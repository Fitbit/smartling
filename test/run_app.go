package test

import (
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
)

func RunApp(app *cli.App, args []string) (string, error) {
	resp, err := CaptureStdout(func() error {
		app.Writer = ioutil.Discard

		return app.Run(args)
	})

	return resp, err
}
