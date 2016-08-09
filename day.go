package gocron

import (
	"fmt"
	"time"
)

func (c *Cron) EveryDays(d int, cmd func()) error {
	if d <= 0 {
		return ErrDay
	}
	c.Schedule(Every(time.Hour*24*time.Duration(d)), FuncJob(cmd))
	return nil
}

func (c *Cron) EveryDay(cmd func()) error {
	c.AddFunc("0 0 0 * * *", cmd)
	return nil
}

func (c *Cron) EveryFixedDay(start int, interval int, cmd func()) error {
	if err := ValidHour(start); err != nil {
		return ErrValueStart
	} else if interval <= 0 {
		return ErrValueInterval
	}
	return c.AddFunc(fmt.Sprintf("0 0 %v/%v * * *", start, interval), cmd)
}

func (c *Cron) EveryHalfDay(cmd func()) error {
	return c.EveryFixedDay(0, 12, cmd)
}

func (c *Cron) EveryQuarterDay(cmd func()) error {
	return c.EveryFixedDay(0, 6, cmd)
}
