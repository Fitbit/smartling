package main

import "os"

func main() {
	err := runApp()

	if err != nil {
		os.Exit(1)
	}
}
