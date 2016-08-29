package main

import (
	"fmt"
	"github.com/fatih/color"
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
		i := 0

		for _, resource := range projectConfig.Resources {
			logInfo(fmt.Sprintf("Using {PathGlob=%v PathExclude=%v}", resource.PathGlob, resource.PathExclude))

			for _, v := range resource.Files() {
				logInfo(color.MagentaString(v))
				i++
			}
		}

		logInfo(fmt.Sprintf("%d files", i))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
