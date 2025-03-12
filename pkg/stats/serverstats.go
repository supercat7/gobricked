package stats

import (
	"fmt"
	"time"
)

var startTime time.Time

func UpTimeInit() {
	startTime = time.Now()
}

func UpTime() time.Duration {
	if startTime.IsZero() {
		startTime = time.Now()
	}
	return time.Since(startTime)
}

func DisplayServerStats() {
	uptime := UpTime()
	fmt.Println("Server Uptime:", uptime.Truncate(time.Second))
}
