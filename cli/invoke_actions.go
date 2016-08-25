package main

import "gopkg.in/urfave/cli.v1"

func invokeActions(actions []action, c *cli.Context) error {
	var err error

	for _, action := range actions {
		err = action(c)

		if err != nil {
			break
		}
	}

	return err
}
