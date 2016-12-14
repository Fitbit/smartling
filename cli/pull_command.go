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
	"github.com/fatih/color"
	"github.com/Fitbit/smartling/model"
	"github.com/Fitbit/smartling/service"
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
			Name:  "retrieval-type",
			Value: "published",
		},
		cli.BoolFlag{
			Name: "include-original-strings",
		},
		cli.IntFlag{
			Name:  "file-uris-limit",
			Value: 20,
		},
	},
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
		retrievalType := c.String("retrieval-type")
		includeOriginalStrings := c.Bool("include-original-strings")
		limit := c.Int("file-uris-limit")
		locales := []string{}

		for locale := range projectConfig.Locales {
			locales = append(locales, locale)
		}

		go func() {
			for _, resource := range projectConfig.Resources {
				for _, files := range resource.LimitFiles(limit) {
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
