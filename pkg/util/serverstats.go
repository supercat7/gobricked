package util

import (
	"gobricked/pkg/global"
	"time"
)

type ServerStats struct {
	uptime time.Duration
}

func UpTimeInit() {
	global.StartTime = time.Now()
}

func UpTime() time.Duration {
	if global.StartTime.IsZero() {
		global.StartTime = time.Now()
	}
	return time.Since(global.StartTime)
}

func GetServerStats() ServerStats {
	uptime := UpTime()
	return ServerStats{
		uptime: uptime.Truncate(time.Second),
	}
}
