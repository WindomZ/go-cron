package cron

import (
	. "github.com/robfig/cron"
	"time"
)

func EveryHours(h int, cmd func()) error {
	if h <= 0 {
		return ErrHour
	}
	getCron().Schedule(Every(time.Hour*time.Duration(h)), FuncJob(cmd))
	return nil
}

func EveryHour(cmd func()) error {
	return EveryHours(1, cmd)
}
