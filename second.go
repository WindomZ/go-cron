package gocron

import "time"

func (c *Cron) EverySeconds(s int, cmd func()) error {
	if s <= 0 {
		return ErrSecond
	}
	c.Schedule(Every(time.Second*time.Duration(s)), FuncJob(cmd))
	return nil
}

func (c *Cron) EverySecond(cmd func()) error {
	return c.EverySeconds(1, cmd)
}
