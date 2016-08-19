package main

import "gopkg.in/urfave/cli.v1"

type action func(c *cli.Context) error
