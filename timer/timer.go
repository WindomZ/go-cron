package timer

import (
	"github.com/WindomZ/go-develop-kit/uuid"
	"sync"
	"time"
)

type Timer struct {
	timers map[string]*time.Timer
	mutex  sync.Mutex
}

func NewTimer() *Timer {
	return &Timer{
		timers: make(map[string]*time.Timer),
	}
}

func (tm *Timer) addTimer(tag string, sec int64, f func()) string {
	if f == nil {
		return tag
	} else if sec <= 0 {
		f()
		return tag
	}
	tm.mutex.Lock()
	if len(tag) == 0 {
		tag = uuid.NewSafeUUID()
	}
	tm.timers[tag] = time.AfterFunc(time.Second*time.Duration(sec), func() {
		if tm != nil {
			tm.removeTimer(tag)
		}
		f()
	})
	tm.mutex.Unlock()
	return tag
}

func (tm *Timer) removeTimer(tag string) *time.Timer {
	if len(tag) != 0 {
		tm.mutex.Lock()
		t, ok := tm.timers[tag]
		if ok {
			delete(tm.timers, tag)
		}
		tm.mutex.Unlock()
		return t
	}
	return nil
}

func (tm *Timer) hasTimer(tag string) bool {
	if len(tag) != 0 {
		tm.mutex.Lock()
		_, ok := tm.timers[tag]
		tm.mutex.Unlock()
		return ok
	}
	return false
}

func (tm *Timer) Close() {
	tm.mutex.Lock()
	for _, t := range tm.timers {
		if t != nil {
			t.Stop()
		}
	}
	tm.mutex.Unlock()
}

func (tm *Timer) AddFunc(sec int64, f func()) string {
	return tm.AddTagFunc("", sec, f)
}

func (tm *Timer) AddTagFunc(tag string, sec int64, f func()) string {
	return tm.addTimer(tag, sec, f)
}

func (tm *Timer) RemoveFunc(tag string) {
	if t := tm.removeTimer(tag); t != nil {
		t.Stop()
	}
}

func (tm *Timer) HasFunc(tag string) bool {
	return tm.hasTimer(tag)
}

func (tm *Timer) IsRunning() bool {
	return len(tm.timers) != 0
}
