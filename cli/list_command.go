package main

import (
	"fmt"
	"github.com/mdreizin/smartling/model"
	"gopkg.in/urfave/cli.v1"
)

var listCommand = cli.Command{
	Name: "list",
	Aliases: []string{
		"ls",
	},
	Usage: "Shows a list of local translations",
	Before: func(c *cli.Context) error {
		return invokeActions([]action{
			ensureMetadataAction,
			injectContainerAction,
			injectProjectConfigAction,
		}, c)
	},
	Action: func(c *cli.Context) error {
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)

		for _, resource := range projectConfig.Resources {
			for _, v := range resource.Files() {
				fmt.Println(v)
			}
		}

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
