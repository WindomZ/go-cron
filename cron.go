package gocron

import (
	"fmt"
	_cron "github.com/robfig/cron"
	"sync"
	"time"
)

type FuncJob _cron.FuncJob

func (f FuncJob) Run() { f() }

type Cron struct {
	_cron.Cron
	Tag     string
	mutex   *sync.Mutex
	running bool
}

func newCron(tags ...string) *Cron {
	var tag string
	if tags != nil && len(tags) != 0 {
		tag = tags[0]
	} else {
		tag = fmt.Sprintf("%v", time.Now().UnixNano())
	}
	return &Cron{
		Cron:    *_cron.New(),
		Tag:     tag,
		mutex:   &sync.Mutex{},
		running: false,
	}
}

func (c *Cron) Start() *Cron {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.running {
		c.Cron.Start()
	}
	c.running = true
	return c
}

func (c *Cron) Stop() *Cron {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Cron.Stop()
	c.running = false
	return c
}

func (c *Cron) IsRunning() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.running
}
