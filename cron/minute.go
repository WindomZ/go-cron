package cron

import (
	"fmt"
	. "github.com/robfig/cron"
	"time"
)

func EveryMinutes(m int, cmd func()) error {
	if m <= 0 {
		return ErrMinute
	}
	getCron().Schedule(Every(time.Minute*time.Duration(m)), FuncJob(cmd))
	return nil
}

func EveryMinute(cmd func()) error {
	return EveryMinutes(1, cmd)
}

func EveryFixedMinute(start int, interval int, cmd func()) error {
	if err := ValidSecond(start); err != nil {
		return ErrValueStart
	} else if err := ValidSecond(interval); err != nil {
		return ErrValueInterval
	} else if err := ValidSecond(start + interval); err != nil {
		return err
	}
	return getCron().AddFunc(fmt.Sprintf("%v/%v * * * * *", start, interval), FuncJob(cmd))
}

func EveryHalfMinute(cmd func()) error {
	return EveryFixedMinute(0, 30, cmd)
}

func EveryQuarterMinute(cmd func()) error {
	return EveryFixedMinute(0, 15, cmd)
}
