package cron

import (
	. "github.com/robfig/cron"
	"time"
)

func EveryDays(d int, cmd func()) error {
	if d <= 0 {
		return ERR_INVALID_DAY
	}
	getCron().Schedule(Every(time.Hour*24*time.Duration(d)), FuncJob(cmd))
	return nil
}

func EveryDay(cmd func()) error {
	getCron().AddFunc("0 0 0 * * *", FuncJob(cmd))
	return nil
}
