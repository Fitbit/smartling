package main

import "gopkg.in/urfave/cli.v1"

func ensureMetadataAction(c *cli.Context) error {
	c.App.Metadata = map[string]interface{}{}

	return nil
}
