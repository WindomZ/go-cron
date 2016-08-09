package cron

import (
	. "github.com/WindomZ/go-cron"
	"sync"
)

var (
	_cron    *Cron
	_mutex   *sync.Mutex = new(sync.Mutex)
	_running bool        = false
)

func init() {
	_cron = NewCron()
}

func StartCron() {
	_mutex.Lock()
	if !_running {
		_cron.Start()
	}
	_running = true
	_mutex.Unlock()
}

func StopCron() {
	_mutex.Lock()
	_cron.Stop()
	_running = false
	_mutex.Unlock()
}

func getCron() *Cron {
	defer StartCron()
	return _cron
}
