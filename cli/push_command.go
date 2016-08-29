package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
	"sync"
	"fmt"
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
		container := c.App.Metadata[containerMetadataKey].(*service.Container)
		authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		wg := &sync.WaitGroup{}

		for _, resource := range projectConfig.Resources {
			directives := resource.Directives.WithPrefix()

			for _, path := range resource.Files() {
				wg.Add(1)

				fmt.Println(path)

				go func(path string, resource model.ProjectResource, directives map[string]string) {
					defer wg.Done()

					container.FileService.Push(&service.FilePushParams{
						ProjectID:  projectConfig.Project.ID,
						FileURI:    projectConfig.FileURI(path),
						FilePath:   path,
						FileType:   resource.Type,
						Authorize:  resource.AuthorizeContent,
						Directives: directives,
						AuthToken:  authToken.AccessToken,
					})
				}(path, resource, directives)
			}
		}

		wg.Wait()

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
