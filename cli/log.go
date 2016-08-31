package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

var logger *log.Logger

func currentLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "", 0)
	}

	logger.SetOutput(os.Stdout)

	return logger
}

func logInfo(v ...interface{}) {
	currentLogger().Println(color.YellowString("[INFO]\t"), fmt.Sprint(v...))
}

func logError(v ...interface{}) {
	currentLogger().Println(color.RedString("[ERROR]\t"), fmt.Sprint(v...))
}
