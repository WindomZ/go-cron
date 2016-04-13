package cron

import (
	. "github.com/robfig/cron"
	"time"
)

func EveryMinutes(m int, cmd func()) error {
	if m <= 0 {
		return ERR_INVALID_MINUTE
	}
	getCron().Schedule(Every(time.Minute*time.Duration(m)), FuncJob(cmd))
	return nil
}

func EveryMinute(cmd func()) error {
	return EveryMinutes(1, cmd)
}
