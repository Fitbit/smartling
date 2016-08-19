package test

import (
	"io/ioutil"
	"os"
	"path"
)

func CopyFile(src string, dst string, f func()) {
	dir := path.Dir(dst)
	bytes, _ := ioutil.ReadFile(src)

	os.MkdirAll(dir, 0777)
	ioutil.WriteFile(dst, bytes, 0644)

	f()

	os.Remove(dst)
}
