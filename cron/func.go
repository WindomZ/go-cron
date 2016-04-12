package cron

import (
	. "github.com/robfig/cron"
	"time"
)

func EverySeconds(s int, cmd func()) error {
	if s <= 0 {
		return ERR_INVALID_SECOND
	}
	getCron().Schedule(Every(time.Second*time.Duration(s)), FuncJob(cmd))
	return nil
}

func EverySecond(cmd func()) error {
	return EverySeconds(1, cmd)
}

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

func EveryHours(h int, cmd func()) error {
	if h <= 0 {
		return ERR_INVALID_HOUR
	}
	getCron().Schedule(Every(time.Hour*time.Duration(h)), FuncJob(cmd))
	return nil
}

func EveryHour(cmd func()) error {
	return EveryHours(1, cmd)
}

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
