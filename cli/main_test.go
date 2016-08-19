package main

import (
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv(nameFor("PROJECT_FILE"), "testdata/smartling.yml")

	os.Exit(m.Run())
}

func TestRun(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	a := assert.New(t)

	os.Args = []string{"cli", "pull"}

	_, err := test.CaptureStdout(func() error {
		main()

		return nil
	})

	a.NoError(err)
}
