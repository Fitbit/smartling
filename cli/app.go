package main

import (
	"gopkg.in/urfave/cli.v1"
	"strings"
)

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "smartling"
	app.Version = strings.TrimPrefix(Version, "v")
	app.Usage = "Smartling CLI to `upload` and `download` translations"
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
