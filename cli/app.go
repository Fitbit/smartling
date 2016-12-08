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
	"gopkg.in/urfave/cli.v1"
	"strings"
)

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "smartling"
	app.Version = strings.TrimPrefix(Version, "v")
	app.Usage = "CLI to work with Smartling translations"
	app.Author = "Marat Dreizin"
	app.Email = "marat.dreizin@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "project-file",
			Value:  ".smartling.yml",
			EnvVar: appKey("project_file"),
		},
		cli.StringFlag{
			Name:   "project-id",
			EnvVar: appKey("project_id"),
		},
		cli.StringFlag{
			Name:   "project-alias",
			EnvVar: appKey("project_alias"),
		},
		cli.StringFlag{
			Name:   "user-id",
			EnvVar: appKey("user_id"),
		},
		cli.StringFlag{
			Name:   "user-secret",
			EnvVar: appKey("user_secret"),
		},
		cli.BoolFlag{
			Name: "no-color",
		},
	}
	app.Before = func(c *cli.Context) error {
		return invokeActions([]action{
			disableColorAction,
		}, c)
	}
	app.Commands = []cli.Command{
		pushCommand,
		pullCommand,
		listCommand,
	}

	return app
}
