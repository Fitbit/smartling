package main

import (
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPullCommand(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	a := assert.New(t)
	app := newApp()
	args := []string{"cli", "--no-color", "pull"}
	resp, err := test.RunApp(app, args)

	a.NoError(err)
	a.NotEqual("", resp)
}
