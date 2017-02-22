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
	"github.com/Fitbit/smartling/logger"
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
			Usage: "Determines the desired format for the download (pending, published, pseudo, contextMatchingInstrumented)",
		},
		cli.BoolFlag{
			Name:  includeOriginalStringsFlagName,
			Usage: "Specifies whether Smartling will return the original string or an empty string where no translation is available",
		},
		cli.IntFlag{
			Name:  fileUrisLimitFlagName,
			Value: 20,
			Usage: "Due to limitation of length of query string it helps to set how many files can be downloaded per one request",
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
					for _, path := range files {
						logger.Infof("%s", color.MagentaString(path))
					}

					batch.Queue(pullWorker(&pullWorkerParams{
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

		for results := range batch.Results() {
			result := results.Value().(*pullWorkerResult)

			if err := results.Error(); err != nil {
				logger.Errorf("[%s] has error %s", color.MagentaString(strings.Join(result.Params.Files, " ")), color.RedString(err.Error()))
			} else {
				for _, file := range result.Files {
					if !visited[file.Path] {
						visited[file.Path] = true
					}
				}

				projectConfig.SaveAllFiles(result.Files, result.Params.Resource)
			}
		}

		logger.Infof("%d files", len(visited))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
