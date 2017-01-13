// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package main

import (
	"github.com/Fitbit/smartling/di"
	"github.com/Fitbit/smartling/model"
	"gopkg.in/urfave/cli.v1"
)

func injectProjectConfigAction(c *cli.Context) error {
	container := c.App.Metadata[containerMetadataKey].(*di.Container)
	project := &model.Project{
		ID:    c.GlobalString(projectIDFlagName),
		Alias: c.GlobalString(projectAliasFlagName),
	}
	userToken := &model.UserToken{
		ID:     c.GlobalString(userTokenIDFlagName),
		Secret: c.GlobalString(userTokenSecretFlagName),
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

	c.App.Metadata[projectConfigMetadataKey] = projectConfig

	return err
}
