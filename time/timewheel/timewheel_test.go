package timewheel_test

import (
	"testing"
	"time"

	"github.com/recallsong/go-utils/time/timewheel"
)

func TestTimeWheel_after(t *testing.T) {
	result, expect := 0, 1
	tw := timewheel.New(time.Millisecond, 0)
	tw.Start()
	tw.After(5*time.Millisecond, func() {
		result++
	})
	time.Sleep(time.Second)
	if result != expect {
		t.Errorf("result is %v, but expect %v", result, expect)
	}
}

func TestTimeWheel_tick(t *testing.T) {
	result, expect := 0, 10
	tw := timewheel.New(time.Millisecond, 0)
	tw.Start()
	tw.Tick(100*time.Millisecond, func() {
		result++
	})
	time.Sleep(time.Second)
	if result != expect {
		t.Errorf("result is %v, but expect %v", result, expect)
	}
}

func TestTimeWheel_afterStop(t *testing.T) {
	result, expect := 0, 0
	tw := timewheel.New(time.Millisecond, 0)
	tw.Start()
	tw.After(100*time.Millisecond, func() {
		result++
	}).Stop()
	time.Sleep(time.Second)
	if result != expect {
		t.Errorf("result is %v, but expect %v", result, expect)
	}
}

func TestTimeWheel_tickStop(t *testing.T) {
	result, expect := 0, 0
	tw := timewheel.New(time.Millisecond, 0)
	tw.Start()
	tw.Tick(100*time.Millisecond, func() {
		result++
	}).Stop()
	time.Sleep(time.Second)
	if result != expect {
		t.Errorf("result is %v, but expect %v", result, expect)
	}
}
