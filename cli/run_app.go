package main

import "os"

func runApp() error {
	app := newApp()

	return app.Run(os.Args)
}
