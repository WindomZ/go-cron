package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestEverySeconds(t *testing.T) {
	defer StopCron()
	EverySeconds(2, func() { fmt.Printf("%v: Hello world!\n", time.Now().Format(time.RFC3339)) })
	select {
	case <-time.After(time.Second*10 + time.Microsecond*10):
		t.Log("Finish")
	}
}
