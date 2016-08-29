package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
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
			Name:  "limit",
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
		defer elapsedTime("Pull", time.Now())

		container := c.App.Metadata[containerMetadataKey].(*service.Container)
		authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		retrievalType := c.String("retrieval-type")
		includeOriginalStrings := c.Bool("include-original-strings")
		limit := c.Int("limit")
		localeIDs := []string{}
		i := 0

		for k := range projectConfig.Locales {
			localeIDs = append(localeIDs, k)
		}

		wg := &sync.WaitGroup{}

		visited := map[string]bool{}

		for _, resource := range projectConfig.Resources {
			for _, f := range resource.PartialFiles(limit) {
				wg.Add(1)

				fileURIs := []string{}

				for _, v := range f {
					fileURIs = append(fileURIs, projectConfig.FileURI(v))
				}

				go func(fileURIs []string, resource model.ProjectResource) {
					defer wg.Done()

					logInfo(fmt.Sprintf("Pull [%s]...", color.MagentaString(strings.Join(fileURIs, " "))))

					files, err := container.FileService.Pull(&service.FilePullParams{
						ProjectID:              projectConfig.Project.ID,
						FileURIs:               fileURIs,
						LocaleIDs:              localeIDs,
						RetrievalType:          retrievalType,
						IncludeOriginalStrings: includeOriginalStrings,
						AuthToken:              authToken.AccessToken,
					})

					if err == nil {
						for _, file := range files {
							if !visited[file.Path] {
								logInfo(fmt.Sprintf("Pulled %s", color.MagentaString(projectConfig.FilePath(file.Path))))

								visited[file.Path] = true
								i++
							}

							projectConfig.SaveFile(file, &resource)
						}
					} else {
						logError(fmt.Sprintf("Pull [%s] was rejected %s", color.MagentaString(strings.Join(fileURIs, " ")), color.RedString(err.Error())))
					}
				}(fileURIs, resource)
			}
		}

		wg.Wait()

		logInfo(fmt.Sprintf("Pulled %d files", i))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
