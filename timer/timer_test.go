package timer

import (
	"testing"
	"time"
)

var _timer *Timer

func TestNewTimer(t *testing.T) {
	_timer = NewTimer()
}

func TestTimer_AddFunc(t *testing.T) {
	sec := time.Now().Unix()
	tag := _timer.AddFunc(1, func() {
		sec = time.Now().Unix() - sec
	})
	if !_timer.hasTimer(tag) {
		t.Fatal("Error AddFunc1")
	}
	time.Sleep(time.Second + time.Millisecond)
	if _timer.hasTimer(tag) {
		t.Fatal("Error AddFunc2")
	} else if sec != 1 {
		t.Fatal("Error AddFunc3", sec)
	}
}

func TestTimer_AddTagFunc(t *testing.T) {
	sec := time.Now().Unix()
	tag := _timer.AddTagFunc("test", 1, func() {
		sec = time.Now().Unix() - sec
	})
	if !_timer.hasTimer(tag) {
		t.Fatal("Error AddTagFunc1")
	}
	time.Sleep(time.Second + time.Millisecond)
	if _timer.hasTimer(tag) {
		t.Fatal("Error AddTagFunc2")
	} else if sec != 1 {
		t.Fatal("Error AddTagFunc3", sec)
	}
}

func TestTimer_RemoveFunc(t *testing.T) {
	tag := _timer.AddFunc(1, func() {
		t.Fatal("Error RemoveFunc")
	})
	time.Sleep(time.Millisecond)
	_timer.RemoveFunc(tag)
}

func TestTimer_IsRunning(t *testing.T) {
	if _timer.IsRunning() {
		t.Fatal("Error IsRunning")
	}
}
