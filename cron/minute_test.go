package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestEveryFixedMinute(t *testing.T) {
	defer StopCron()
	EveryFixedMinute(0, 5, func() { fmt.Printf("%v: EveryFixedMinute\n", time.Now().Format(time.RFC3339)) })
	EveryHalfMinute(func() { fmt.Printf("%v: EveryHalfMinute\n", time.Now().Format(time.RFC3339)) })
	EveryQuarterMinute(func() { fmt.Printf("%v: EveryQuarterMinute\n", time.Now().Format(time.RFC3339)) })
	select {
	case <-time.After(time.Minute + time.Microsecond*10):
		t.Log("Finish")
	}
}
