package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
	"sync"
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
		defer elapsedTime("Push", time.Now())

		container := c.App.Metadata[containerMetadataKey].(*service.Container)
		authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		i := 0
		wg := &sync.WaitGroup{}

		for _, resource := range projectConfig.Resources {
			directives := resource.Directives.WithPrefix()

			for _, path := range resource.Files() {
				wg.Add(1)

				logInfo(fmt.Sprintf("Push %s...", color.MagentaString(path)))

				go func(path string, resource model.ProjectResource, directives map[string]string) {
					defer wg.Done()

					stats, err := container.FileService.Push(&service.FilePushParams{
						ProjectID:  projectConfig.Project.ID,
						FileURI:    projectConfig.FileURI(path),
						FilePath:   path,
						FileType:   resource.Type,
						Authorize:  resource.AuthorizeContent,
						Directives: directives,
						AuthToken:  authToken.AccessToken,
					})

					if err == nil {
						logInfo(fmt.Sprintf("Pushed %s {Override=%t Strings=%d Words=%d}", color.MagentaString(path), stats.OverWritten, stats.StringCount, stats.WordCount))
						i++
					} else {
						logError(fmt.Sprintf("Push %s was rejected %s", path, color.RedString(err.Error())))
					}
				}(path, resource, directives)
			}
		}

		wg.Wait()

		logInfo(fmt.Sprintf("Pushed %d files", i))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
