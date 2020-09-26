package emitter

import (
	"fmt"
	"math"
	"sync"
)

// MaxListenersPerEvent .
const MaxListenersPerEvent = math.MaxUint32

// Event .
type Event struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// Listener .
type Listener func(e *Event)

// Watcher .
type Watcher interface {
	Channel(size int) (<-chan *Event, error)
	Callback(listener Listener, first ...func(string) error) error
	Close() error
	Stop(last ...func(name string))
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
		key:     MaxListenersPerEvent,
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

func (e *Emitter) appendListener(name string, ln Listener, first func(name string) error) (uint64, error) {
	e.lock.Lock()
	listeners, ok := e.listeners[name]
	if ok {
		if len(listeners) >= MaxListenersPerEvent-1 {
			e.lock.Unlock()
			return 0, fmt.Errorf("large map size in Emitter.listeners")
			// panic("large map size in Emitter.listeners")
		}
	} else {
		if first != nil {
			err := first(name)
			if err != nil {
				e.lock.Unlock()
				return 0, err
			}
		}
		listeners = make(map[uint64]Listener)
		e.listeners[name] = listeners
	}
	id := e.index
	for {
		if _, ok := listeners[id]; ok {
			e.index++
			if e.index == MaxListenersPerEvent {
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
	return id, nil
}

func (e *Emitter) removeListener(name string, key uint64, last func(name string)) {
	e.lock.Lock()
	listeners, ok := e.listeners[name]
	if ok {
		if _, ok := listeners[key]; ok {
			delete(listeners, key)
			if len(listeners) <= 0 {
				if last != nil {
					last(name)
				}
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

func (w *watcher) Channel(size int) (<-chan *Event, error) {
	ch := make(chan *Event, size)
	err := w.Callback(Listener(func(e *Event) {
		ch <- e
	}))
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (w *watcher) Callback(listener Listener, first ...func(string) error) error {
	var f func(string) error
	if len(first) > 0 {
		f = first[0]
	}
	key, err := w.emitter.appendListener(w.name, listener, f)
	if err != nil {
		return err
	}
	w.key = key
	return nil
}

func (w *watcher) Close() error {
	if w.key == MaxListenersPerEvent {
		return nil
	}
	w.emitter.removeListener(w.name, w.key, nil)
	w.key = MaxListenersPerEvent
	return nil
}

func (w *watcher) Stop(last ...func(string)) {
	if w.key == MaxListenersPerEvent {
		return
	}
	var l func(string)
	if len(last) > 0 {
		l = last[0]
	}
	w.emitter.removeListener(w.name, w.key, l)
	w.key = MaxListenersPerEvent
}
