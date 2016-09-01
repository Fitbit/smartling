package main

import (
	"errors"
	"github.com/mdreizin/smartling/model"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

func validateProjectConfigAction(c *cli.Context) error {
	var err error

	issues := []string{}

	projectConfig := c.App.Metadata[projectConfigKey].(*model.ProjectConfig)

	if projectConfig.Project.ID == "" {
		issues = append(issues, "project-id is required")
	}

	if projectConfig.UserToken.ID == "" {
		issues = append(issues, "user-id is required")
	}

	if projectConfig.UserToken.Secret == "" {
		issues = append(issues, "user-secret is required")
	}

	if len(issues) > 0 {
		err = errors.New(strings.Join(issues, "\n"))
	}

	return err
}
