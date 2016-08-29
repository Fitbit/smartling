package main

import (
	"github.com/fatih/color"
	"log"
)

func logInfo(v ...interface{}) {
	log.SetPrefix(color.YellowString("[INFO]\t"))
	log.SetFlags(0)
	log.Println(v...)
}

func logError(v ...interface{}) {
	log.SetPrefix(color.RedString("[ERROR]\t"))
	log.SetFlags(0)
	log.Println(v...)
}
