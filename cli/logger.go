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

func logPrefix(level string) string {
	return fmt.Sprintf("[%s]", level)
}

func logInfo(v ...interface{}) {
	currentLogger().Println(color.YellowString(logPrefix(logLevelInfo)), fmt.Sprint(v...))
}

func logError(v ...interface{}) {
	currentLogger().Println(color.RedString(logPrefix(logLevelError)), fmt.Sprint(v...))
}
