package cron

import (
	. "github.com/robfig/cron"
	"sync"
)

var (
	_cron    *Cron
	_mutex   *sync.Mutex
	_running bool
)

func init() {
	_cron = New()
	_mutex = &sync.Mutex{}
	_running = false
}

func StartCron() {
	_mutex.Lock()
	defer _mutex.Unlock()
	if !_running {
		_cron.Start()
	}
	_running = true
}

func StopCron() {
	_mutex.Lock()
	defer _mutex.Unlock()
	_cron.Stop()
	_running = false
}

func getCron() *Cron {
	defer StartCron()
	return _cron
}
