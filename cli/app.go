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

func init() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show help",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Output the version number",
	}
}

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "smartling"
	app.Version = strings.TrimPrefix(Version, "v")
	app.Usage = "CLI to work with Smartling translations"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   projectFileFlagName,
			Value:  ".smartling.yml",
			EnvVar: envVar(projectFileFlagName),
			Usage:  "Project configuration file",
		},
		cli.StringFlag{
			Name:   projectIDFlagName,
			EnvVar: envVar(projectIDFlagName),
			Usage:  "Project identifier for your Smartling v2 API Token",
		},
		cli.StringFlag{
			Name:   projectAliasFlagName,
			EnvVar: envVar(projectAliasFlagName),
			Usage:  "Unique alias of your project",
		},
		cli.StringFlag{
			Name:   userTokenIDFlagName,
			EnvVar: envVar(userTokenIDFlagName),
			Usage:  "User identifier for your Smartling v2 API Token",
		},
		cli.StringFlag{
			Name:   userTokenSecretFlagName,
			EnvVar: envVar(userTokenSecretFlagName),
			Usage:  "Token secret for your Smartling v2 API Token",
		},
		cli.BoolFlag{
			Name:  noColorFlagName,
			Usage: "Turn off colored output for log messages (default: false)",
		},
		cli.BoolFlag{
			Name:  verboseFlagName,
			Usage: "Output verbose messages on internal operations (default: false)",
		},
		cli.BoolTFlag{
			Name:  saveAccessTokenFlagName,
			Usage: "Save Smartling v2 API Access Token (default: true)",
		},
	}
	app.Before = func(c *cli.Context) error {
		return invokeActions([]action{
			disableColorsAction,
		}, c)
	}
	app.Commands = []cli.Command{
		pushCommand,
		pullCommand,
		listCommand,
	}

	return app
}
