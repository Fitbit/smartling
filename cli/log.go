package main

import (
	"fmt"
	"github.com/fatih/color"
)

func logInfo(v ...interface{}) {
	fmt.Print(color.YellowString("[INFO]\t"))
	fmt.Println(v...)
}

func logError(v ...interface{}) {
	fmt.Print(color.RedString("[ERROR]\t"))
	fmt.Println(v...)
}
