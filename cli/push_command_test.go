package main

import (
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushCommand(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	a := assert.New(t)
	app := newApp()
	args := []string{"cli", "push"}
	resp, err := test.RunApp(app, args)

	a.NoError(err)
	a.NotEqual("", resp)
}
