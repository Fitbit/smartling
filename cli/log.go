package main

import (
	"fmt"
	"github.com/fatih/color"
)

func logInfo(v ...interface{}) {
	fmt.Println(color.YellowString("[INFO]\t"), fmt.Sprint(v...))
}

func logError(v ...interface{}) {
	fmt.Println(color.RedString("[ERROR]\t%v"), fmt.Sprint(v...))
}
