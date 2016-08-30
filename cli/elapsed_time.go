package main

import (
	"time"
)

func elapsedTime(startTime time.Time) {
	endTime := time.Since(startTime)

	logInfo(endTime)
}
