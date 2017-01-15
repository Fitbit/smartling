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
	"fmt"
	"github.com/Fitbit/smartling/di"
	"github.com/Fitbit/smartling/model"
	"github.com/fatih/color"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/urfave/cli.v1"
	"runtime"
	"strings"
	"time"
)

var pullCommand = cli.Command{
	Name:  "pull",
	Usage: "Downloads translations",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  retrievalTypeFlagName,
			Value: "published",
		},
		cli.BoolFlag{
			Name: includeOriginalStringsFlagName,
		},
		cli.IntFlag{
			Name:  fileUrisLimitFlagName,
			Value: 20,
		},
	},
	Before: func(c *cli.Context) error {
		return invokeActions([]action{
			injectDiContainerAction,
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

		container := c.App.Metadata[containerMetadataKey].(*di.Container)
		authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		retrievalType := c.String(retrievalTypeFlagName)
		includeOriginalStrings := c.Bool(includeOriginalStringsFlagName)
		limit := c.Int(fileUrisLimitFlagName)
		locales := []string{}

		for locale := range projectConfig.Locales {
			locales = append(locales, locale)
		}

		go func() {
			for _, resource := range projectConfig.Resources {
				for _, files := range resource.BatchFiles(limit) {
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

		for result := range batch.Results() {
			resp := result.Value().(*pullResponse)

			if err := result.Error(); err != nil {
				logError(fmt.Sprintf("[%s] has error %s", color.MagentaString(strings.Join(resp.Request.Files, " ")), color.RedString(err.Error())))
			} else {
				for _, file := range resp.Files {
					if !visited[file.Path] {
						visited[file.Path] = true
					}
				}

				projectConfig.SaveAllFiles(resp.Files, resp.Request.Resource)
			}
		}

		logInfo(fmt.Sprintf("%d files", len(visited)))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
