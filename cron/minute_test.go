package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestEveryFixedMinute(t *testing.T) {
	defer StopCron()
	if err := EveryFixedMinute(0, 5, func() { fmt.Printf("%v: 0/5\n", time.Now().Format(time.RFC3339)) }); err != nil {
		t.Fatal(err)
	}
	if err := EveryFixedMinute(0, 65, func() { fmt.Printf("%v: 0/65\n", time.Now().Format(time.RFC3339)) }); err != nil {
		t.Fatal(err)
	}
	if err := EveryHalfMinute(func() { fmt.Printf("%v: 0/30\n", time.Now().Format(time.RFC3339)) }); err != nil {
		t.Fatal(err)
	}
	if err := EveryQuarterMinute(func() { fmt.Printf("%v: 0/15\n", time.Now().Format(time.RFC3339)) }); err != nil {
		t.Fatal(err)
	}
	select {
	case <-time.After(time.Minute*2 + time.Microsecond*10):
		t.Log("Finish")
	}
}
