// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
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
