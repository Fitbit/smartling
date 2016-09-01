package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
)

func persistAuthTokenAction(c *cli.Context) (err error) {
	if c.App.Metadata[containerKey] != nil {
		container := c.App.Metadata[containerKey].(*service.Container)

		if c.App.Metadata[authTokenKey] != nil {
			authToken := c.App.Metadata[authTokenKey].(*model.AuthToken)
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
