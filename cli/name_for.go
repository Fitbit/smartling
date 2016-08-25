package main

import "strings"

func nameFor(name string) string {
	return "SMARTLING_" + strings.ToUpper(name)
}
