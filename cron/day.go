package cron

import (
	"fmt"
	. "github.com/robfig/cron"
	"time"
)

func EveryDays(d int, cmd func()) error {
	if d <= 0 {
		return ErrDay
	}
	getCron().Schedule(Every(time.Hour*24*time.Duration(d)), FuncJob(cmd))
	return nil
}

func EveryDay(cmd func()) error {
	getCron().AddFunc("0 0 0 * * *", FuncJob(cmd))
	return nil
}

func EveryFixedDay(start int, interval int, cmd func()) error {
	if err := ValidHour(start); err != nil {
		return ErrValueStart
	} else if err := ValidHour(interval); err != nil {
		return ErrValueInterval
	} else if err := ValidHour(start + interval); err != nil {
		return err
	}
	return getCron().AddFunc(fmt.Sprintf("* * %v/%v * * *", start, interval), FuncJob(cmd))
}

func EveryHalfDay(cmd func()) error {
	return EveryFixedDay(0, 30, cmd)
}

func EveryQuarterDay(cmd func()) error {
	return EveryFixedDay(0, 15, cmd)
}
