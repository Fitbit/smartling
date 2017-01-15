// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package logger

import (
	"fmt"
	"github.com/fatih/color"
)

func Info(v ...interface{}) {
	current().Println(color.YellowString(prefix(infoLevel)), fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	current().Println(color.RedString(prefix(errorLevel)), fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	Error(fmt.Sprintf(format, v...))
}

func DisableColors(colors bool) {
	color.NoColor = colors
}
