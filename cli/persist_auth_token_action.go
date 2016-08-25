package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
)

func persistAuthTokenAction(c *cli.Context) (err error) {
	if c.App.Metadata[containerMetadataKey] != nil {
		container := c.App.Metadata[containerMetadataKey].(*service.Container)

		if c.App.Metadata[authTokenMetadataKey] != nil {
			authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
			projectConfig := model.ProjectConfig{
				AuthToken: model.AuthToken{
					AccessToken: authToken.RefreshToken,
				},
			}

			err = container.ProjectConfigService.UpdateConfig(&projectConfig)
		}
	}

	return err
}
