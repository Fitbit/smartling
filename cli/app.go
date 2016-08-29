package main

import (
	"gopkg.in/urfave/cli.v1"
)

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "smartling"
	app.Version = Version
	app.Usage = "Smartling CLI to `upload` and `download` translations"
	app.Author = "Marat Dreizin"
	app.Email = "marat.dreizin@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "project-file",
			Value:  ".smartling.yml",
			EnvVar: nameFor("project_file"),
		},
		cli.StringFlag{
			Name:   "project-id",
			EnvVar: nameFor("project_id"),
		},
		cli.StringFlag{
			Name:   "project-alias",
			EnvVar: nameFor("project_alias"),
		},
		cli.StringFlag{
			Name:   "user-id",
			EnvVar: nameFor("user_id"),
		},
		cli.StringFlag{
			Name:   "user-secret",
			EnvVar: nameFor("user_secret"),
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
