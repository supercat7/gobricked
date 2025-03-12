package stats

import (
	"time"
)

type ServerStats struct {
	uptime time.Duration
}

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

func GetServerStats() ServerStats {
	uptime := UpTime()
	return ServerStats{
		uptime: uptime.Truncate(time.Second),
	}
}
