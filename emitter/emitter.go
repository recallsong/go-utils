package emitter

import (
	"fmt"
	"math"
	"sync"
)

// Event .
type Event struct {
	Name string
	Data interface{}
}

// Listener .
type Listener func(e *Event)

// Watcher .
type Watcher interface {
	Close() error
	Channel(size int) <-chan *Event
	Callback(listener Listener)
}

// Emitter .
type Emitter struct {
	index     uint64
	listeners map[string]map[uint64]Listener
	lock      sync.RWMutex
}

// New .
func New() *Emitter {
	return &Emitter{
		listeners: make(map[string]map[uint64]Listener),
	}
}

// Watch .
func (e *Emitter) Watch(name string) Watcher {
	return &watcher{
		name:    name,
		emitter: e,
		key:     math.MaxUint64,
	}
}

// Emit .
func (e *Emitter) Emit(event *Event) int {
	e.lock.RLock()
	var num int
	listeners, ok := e.listeners[event.Name]
	if ok {
		for _, ln := range listeners {
			ln(event)
			num++
		}
	}
	e.lock.RUnlock()
	return num
}

func (e *Emitter) String() string {
	return fmt.Sprint(e.index, e.listeners)
}

func (e *Emitter) appendListener(name string, ln Listener) uint64 {
	e.lock.Lock()
	listeners, ok := e.listeners[name]
	if ok {
		if len(listeners) >= math.MaxUint16-1 {
			e.lock.Unlock()
			panic("large map size in Emitter.listeners")
		}
	} else {
		listeners = make(map[uint64]Listener)
		e.listeners[name] = listeners
	}
	id := e.index
	for {
		if _, ok := listeners[id]; ok {
			e.index++
			if e.index == math.MaxUint64 {
				e.index = 0
			}
			id = e.index
		} else {
			break
		}
	}
	listeners[id] = ln
	e.index++
	e.lock.Unlock()
	return id
}

func (e *Emitter) removeListener(name string, key uint64) {
	e.lock.Lock()
	listeners, ok := e.listeners[name]
	if ok {
		if _, ok := listeners[key]; ok {
			delete(listeners, key)
			if len(listeners) <= 0 {
				delete(e.listeners, name)
			}
		}
	}
	e.lock.Unlock()
}

// watcher .
type watcher struct {
	emitter *Emitter
	name    string
	key     uint64
}

func (w *watcher) Channel(size int) <-chan *Event {
	ch := make(chan *Event, size)
	w.Callback(Listener(func(e *Event) {
		ch <- e
	}))
	return ch
}

func (w *watcher) Callback(listener Listener) {
	w.key = w.emitter.appendListener(w.name, listener)
}

func (w *watcher) Close() error {
	if w.key == math.MaxUint64 {
		return nil
	}
	w.emitter.removeListener(w.name, w.key)
	return nil
}
