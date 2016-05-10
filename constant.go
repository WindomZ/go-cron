package gocron

import (
	_cron "github.com/robfig/cron"
	"time"
)

func Every(duration time.Duration) _cron.ConstantDelaySchedule {
	return _cron.Every(duration)
}
