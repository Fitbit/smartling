package main

import (
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
	"path/filepath"
)

func injectContainerAction(c *cli.Context) error {
	container := service.Container{}

	filename, err := filepath.Abs(c.GlobalString("project-file"))

	if err == nil {
		err = container.SetUp(filename)
	}

	c.App.Metadata[containerMetadataKey] = &container

	return err
}
