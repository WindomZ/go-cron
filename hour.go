package gocron

import (
	"fmt"
	"time"
)

func (c *Cron) EveryHours(h int, cmd func()) error {
	if h <= 0 {
		return ErrHour
	}
	c.Schedule(Every(time.Hour*time.Duration(h)), FuncJob(cmd))
	return nil
}

func (c *Cron) EveryHour(cmd func()) error {
	return c.EveryHours(1, cmd)
}

func (c *Cron) EveryFixedHour(start int, interval int, cmd func()) error {
	if err := ValidMinute(start); err != nil {
		return ErrValueStart
	} else if interval <= 0 {
		return ErrValueInterval
	}
	return c.AddFunc(fmt.Sprintf("0 %v/%v * * * *", start, interval), FuncJob(cmd))
}

func (c *Cron) EveryHalfHour(cmd func()) error {
	return c.EveryFixedHour(0, 30, cmd)
}

func (c *Cron) EveryQuarterHour(cmd func()) error {
	return c.EveryFixedHour(0, 15, cmd)
}
