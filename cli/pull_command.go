package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/urfave/cli.v1"
	"runtime"
	"strings"
	"sync"
	"time"
)

var pullCommand = cli.Command{
	Name:  "pull",
	Usage: "Downloads translations",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "retrieval-type",
			Value: "published",
		},
		cli.BoolFlag{
			Name: "include-original-strings",
		},
		cli.IntFlag{
			Name:  "file-uris-limit",
			Value: 20,
		},
	},
	Before: func(c *cli.Context) error {
		return invokeActions([]action{
			ensureMetadataAction,
			injectContainerAction,
			injectProjectConfigAction,
			validateProjectConfigAction,
			injectAuthTokenAction,
		}, c)
	},
	Action: func(c *cli.Context) error {
		defer elapsedTime(time.Now())

		p := pool.NewLimited(uint(runtime.NumCPU()))

		defer p.Close()

		batch := p.Batch()

		container := c.App.Metadata[containerMetadataKey].(*service.Container)
		authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		retrievalType := c.String("retrieval-type")
		includeOriginalStrings := c.Bool("include-original-strings")
		limit := c.Int("file-uris-limit")
		locales := []string{}

		for locale := range projectConfig.Locales {
			locales = append(locales, locale)
		}

		go func() {
			for _, resource := range projectConfig.Resources {
				for _, files := range resource.LimitFiles(limit) {
					batch.Queue(pullJob(&pullRequest{
						Files:                  files,
						Locales:                locales,
						IncludeOriginalStrings: includeOriginalStrings,
						RetrievalType:          retrievalType,
						Config:                 projectConfig,
						Resource:               &resource,
						AuthToken:              authToken.AccessToken,
						FileService:            container.FileService,
					}))
				}
			}

			batch.QueueComplete()
		}()

		visited := map[string]bool{}

		wg := &sync.WaitGroup{}

		for result := range batch.Results() {
			resp := result.Value().(*pullResponse)

			if err := result.Error(); err != nil {
				logError(fmt.Sprintf("[%s] have error %s", color.MagentaString(strings.Join(resp.Request.Files, " ")), color.RedString(err.Error())))
			} else {
				for _, file := range resp.Files {
					wg.Add(1)

					if !visited[file.Path] {
						visited[file.Path] = true
					}

					go func(file *model.File, resource *model.ProjectResource) {
						defer wg.Done()

						projectConfig.SaveFile(file, resource)
					}(file, resp.Request.Resource)
				}
			}
		}

		wg.Wait()

		logInfo(fmt.Sprintf("%d files", len(visited)))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
