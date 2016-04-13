package cron

import (
	"fmt"
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

func EveryFixedHour(start int, interval int, cmd func()) error {
	if err := ValidMinute(start); err != nil {
		return ErrValueStart
	} else if interval <= 0 {
		return ErrValueInterval
	}
	return getCron().AddFunc(fmt.Sprintf("* %v/%v * * * *", start, interval), FuncJob(cmd))
}

func EveryHalfHour(cmd func()) error {
	return EveryFixedHour(0, 30, cmd)
}

func EveryQuarterHour(cmd func()) error {
	return EveryFixedHour(0, 15, cmd)
}
