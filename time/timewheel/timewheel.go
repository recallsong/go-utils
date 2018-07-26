package timewheel

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"github.com/recallsong/go-utils/lang"
)

type Timer interface {
	Stop()
}

type timer struct {
	_      lang.NoCopy
	d      time.Duration
	circle int
	pos    int
	tw     *TimeWheel
	fn     func()
	once   bool
}

func (t *timer) Stop() {
	t.tw.delCh <- t
}

type TimeWheel struct {
	_        lang.NoCopy
	interval time.Duration
	ticker   *time.Ticker
	slots    []*list.List
	slotSize int
	position int
	addCh    chan *timer
	delCh    chan *timer
	stopCh   chan struct{}
	once     sync.Once
}

func New(interval time.Duration, slotSize int) *TimeWheel {
	if interval <= 0 {
		panic("invalid timewheel interval")
	}
	if slotSize <= 0 {
		slotSize = 3600
	}
	slots := make([]*list.List, slotSize)
	for i := 0; i < slotSize; i++ {
		slots[i] = list.New()
	}
	return &TimeWheel{
		interval: interval,
		slots:    slots,
		slotSize: slotSize,
		position: 0,
		addCh:    make(chan *timer),
		delCh:    make(chan *timer),
		stopCh:   make(chan struct{}),
	}
}

func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	go func(tw *TimeWheel) {
		for {
			select {
			case <-tw.ticker.C:
				tw.doTick()
			case t := <-tw.addCh:
				tw.doAddTimer(t)
			case t := <-tw.delCh:
				tw.doDeleteTimer(t)
			case <-tw.stopCh:
				tw.ticker.Stop()
				return
			}
		}
	}(tw)
}

func (tw *TimeWheel) doTick() {
	list := tw.slots[tw.position]
	for e := list.Front(); e != nil; {
		t := e.Value.(*timer)
		if t.circle > 0 {
			t.circle--
			e = e.Next()
			continue
		}
		next := e.Next()
		list.Remove(e)
		e = next
		fmt.Println(".", t.once)
		if !t.once {

			tw.doAddTimer(t)
		}
		go t.fn()
	}
	tw.position = (tw.position + 1) % tw.slotSize
}

func (tw *TimeWheel) doAddTimer(t *timer) {
	ticks := int(t.d / tw.interval)
	t.circle = int(ticks) / tw.slotSize
	t.pos = (tw.position + ticks) % tw.slotSize
	tw.slots[t.pos].PushBack(t)
}

func (tw *TimeWheel) doDeleteTimer(t *timer) {
	if t.pos > 0 && t.pos < tw.slotSize {
		list := tw.slots[t.pos]
		for e := list.Front(); e != nil; {
			val := e.Value.(*timer)
			if val == t {
				list.Remove(e)
				break
			}
			e = e.Next()
		}
	}
}

func (tw *TimeWheel) Stop() {
	tw.once.Do(func() {
		close(tw.stopCh)
	})
}

func (tw *TimeWheel) Tick(d time.Duration, fn func()) Timer {
	return tw.addTimer(d, fn, false)
}

func (tw *TimeWheel) After(d time.Duration, fn func()) Timer {
	return tw.addTimer(d, fn, true)
}

func (tw *TimeWheel) addTimer(d time.Duration, fn func(), once bool) *timer {
	if int64(d) <= 0 {
		panic(fmt.Sprintf("timewheel tick argement invalid d=%v", int64(d)))
	}
	if fn == nil {
		panic("timewheel tick argement fn should not be nil")
	}
	t := &timer{d: d, tw: tw, fn: fn, once: once}
	tw.addCh <- t
	return t
}
