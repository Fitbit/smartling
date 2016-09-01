package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
)

func injectProjectConfigAction(c *cli.Context) error {
	container := c.App.Metadata[containerKey].(*service.Container)
	project := &model.Project{
		ID:    c.GlobalString("project-id"),
		Alias: c.GlobalString("project-alias"),
	}
	userToken := &model.UserToken{
		ID:     c.GlobalString("user-id"),
		Secret: c.GlobalString("user-secret"),
	}
	projectConfig, err := container.ProjectConfigService.GetConfig()

	if err == nil {
		if project.ID != "" {
			projectConfig.Merge(&model.ProjectConfig{
				Project: model.Project{
					ID: project.ID,
				},
			})
		}

		if project.Alias != "" {
			projectConfig.Merge(&model.ProjectConfig{
				Project: model.Project{
					Alias: project.Alias,
				},
			})
		}

		if userToken.ID != "" {
			projectConfig.Merge(&model.ProjectConfig{
				UserToken: model.UserToken{
					ID: userToken.ID,
				},
			})
		}

		if userToken.Secret != "" {
			projectConfig.Merge(&model.ProjectConfig{
				UserToken: model.UserToken{
					Secret: userToken.Secret,
				},
			})
		}
	}

	c.App.Metadata[projectConfigKey] = projectConfig

	return err
}
