package main

import (
	"github.com/fatih/color"
	"fmt"
)

func logInfo(v ...interface{}) {
	fmt.Print(color.YellowString("[INFO]\t"))
	fmt.Println(v...)
}

func logError(v ...interface{}) {
	fmt.Print(color.RedString("[ERROR]\t"))
	fmt.Println(v...)
}
