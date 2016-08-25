package test

import (
	"bytes"
	"io"
	"os"
)

func CaptureStdout(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer

	io.Copy(&buf, r)

	return buf.String(), err
}
