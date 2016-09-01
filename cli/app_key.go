package main

import "strings"

func appKey(name string) string {
	return "SMARTLING_" + strings.ToUpper(name)
}
