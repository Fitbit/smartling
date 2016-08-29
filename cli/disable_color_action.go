package main

import (
	"github.com/fatih/color"
	"gopkg.in/urfave/cli.v1"
)

func disableColorAction(c *cli.Context) error {
	color.NoColor = c.GlobalBool("no-color")

	return nil
}
