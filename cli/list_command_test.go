package main

import (
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListCommand(t *testing.T) {
	a := assert.New(t)
	app := newApp()
	args := []string{"cli", "list"}
	resp, err := test.RunApp(app, args)

	a.NoError(err)
	a.NotEqual("", resp)
}
