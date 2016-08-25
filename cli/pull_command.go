package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
	"sync"
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
	},
	Before: func(c *cli.Context) error {
		return invokeActions(globalBeforeActions, c)
	},
	Action: func(c *cli.Context) error {
		container := c.App.Metadata[containerMetadataKey].(*service.Container)
		authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		retrievalType := c.String("retrieval-type")
		includeOriginalStrings := c.Bool("include-original-strings")
		localeIDs := []string{}

		for k := range projectConfig.Locales {
			localeIDs = append(localeIDs, k)
		}

		wg := &sync.WaitGroup{}

		for _, resource := range projectConfig.Resources {
			wg.Add(1)

			fileURIs := []string{}

			for _, v := range resource.Files() {
				fileURIs = append(fileURIs, projectConfig.FileURI(v))
			}

			go func(fileURIs []string, resource model.ProjectResource) {
				defer wg.Done()

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
						projectConfig.SaveFile(file, &resource)
					}
				}
			}(fileURIs, resource)
		}

		wg.Wait()

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions(globalAfterActions, c)
	},
}
