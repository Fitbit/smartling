package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/urfave/cli.v1"
	"runtime"
	"time"
)

var pushCommand = cli.Command{
	Name:  "push",
	Usage: "Uploads translations",
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

		container := c.App.Metadata[containerKey].(*service.Container)
		authToken := c.App.Metadata[authTokenKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigKey].(*model.ProjectConfig)

		go func() {
			for _, resource := range projectConfig.Resources {
				for _, path := range resource.Files() {
					batch.Queue(pushJob(&pushRequest{
						Path:        path,
						Resource:    &resource,
						Config:      projectConfig,
						AuthToken:   authToken.AccessToken,
						FileService: container.FileService,
					}))
				}
			}

			batch.QueueComplete()
		}()

		i := 0

		for result := range batch.Results() {
			resp := result.Value().(*pushResponse)

			if err := result.Error(); err != nil {
				logError(fmt.Sprintf("%s has error %s", resp.Params.FilePath, color.RedString(err.Error())))
			} else {
				logInfo(fmt.Sprintf("%s {Override=%t Strings=%d Words=%d}", color.MagentaString(resp.Params.FilePath), resp.Stats.OverWritten, resp.Stats.StringCount, resp.Stats.WordCount))
				i++
			}
		}

		logInfo(fmt.Sprintf("%d files", i))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
