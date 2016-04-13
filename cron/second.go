package cron

import (
	. "github.com/robfig/cron"
	"time"
)

func EverySeconds(s int, cmd func()) error {
	if s <= 0 {
		return ErrSecond
	}
	getCron().Schedule(Every(time.Second*time.Duration(s)), FuncJob(cmd))
	return nil
}

func EverySecond(cmd func()) error {
	return EverySeconds(1, cmd)
}
