package gocron

import (
	"fmt"
	"time"
)

func (c *Cron) EveryMinutes(m int, cmd func()) error {
	if m <= 0 {
		return ErrMinute
	}
	c.Schedule(Every(time.Minute*time.Duration(m)), FuncJob(cmd))
	return nil
}

func (c *Cron) EveryMinute(cmd func()) error {
	return c.EveryMinutes(1, cmd)
}

func (c *Cron) EveryFixedMinute(start int, interval int, cmd func()) error {
	if err := ValidSecond(start); err != nil {
		return ErrValueStart
	} else if interval <= 0 {
		return ErrValueInterval
	}
	return c.AddFunc(fmt.Sprintf("%v/%v * * * * *", start, interval), FuncJob(cmd))
}

func (c *Cron) EveryHalfMinute(cmd func()) error {
	return c.EveryFixedMinute(0, 30, cmd)
}

func (c *Cron) EveryQuarterMinute(cmd func()) error {
	return c.EveryFixedMinute(0, 15, cmd)
}
