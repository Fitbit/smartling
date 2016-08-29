package main

import (
	"fmt"
	"time"
)

func elapsedTime(name string, startTime time.Time) {
	endTime := time.Since(startTime)

	logInfo(fmt.Sprintf("%s took", name), endTime)
}
