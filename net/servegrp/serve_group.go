package servegrp

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/recallsong/go-utils/errorx"
)

type ServeItem interface {
	Serve() error
	Close() error
}

// ServeGroup 管理多个ServeItem的启动和关闭
type ServeGroup struct {
	serves map[string]ServeItem
	lock   sync.RWMutex
	wg     sync.WaitGroup
}

func NewServeGroup() *ServeGroup {
	return &ServeGroup{serves: make(map[string]ServeItem)}
}

func (sg *ServeGroup) Put(key string, s ServeItem) error {
	sg.lock.Lock()
	defer sg.lock.Unlock()
	if _, ok := sg.serves[key]; ok {
		return fmt.Errorf("%s is exist already", key)
	}
	sg.serves[key] = s
	return nil
}

func (sg *ServeGroup) Num() int {
	sg.lock.RLock()
	defer sg.lock.RUnlock()
	return len(sg.serves)
}

func (sg *ServeGroup) Serve(closeCh <-chan os.Signal, stopFn func(err error, key string, svr ServeItem)) error {
	sg.lock.RLock()
	num := len(sg.serves)
	ch := make(chan error, num)
	sg.wg.Add(num)
	for key, svr := range sg.serves {
		go func(key string, svr ServeItem) {
			var err error
			defer func() {
				sg.lock.Lock()
				if sg.serves[key] == svr {
					delete(sg.serves, key)
				}
				sg.lock.Unlock()
				sg.wg.Done()
				stopFn(err, key, svr)
				ch <- err
			}()
			err = svr.Serve()
		}(key, svr)
	}
	sg.lock.RUnlock()
	runtime.Gosched() // let serve goroutine work

	// wait to stop
	errs := errorx.Errors{}
	for i := 0; i < num; i++ {
		select {
		case <-closeCh:
			return sg.Close()
		case err := <-ch:
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs.MaybeUnwrap()
}

func (sg *ServeGroup) CloseOne(addr string) error {
	sg.lock.Lock()
	defer sg.lock.Unlock()
	if svr, ok := sg.serves[addr]; ok {
		err := svr.Close()
		if err != nil {
			return err
		}
		delete(sg.serves, addr)
	}
	return nil
}

// Close 关闭所有的ServeItem
func (sg *ServeGroup) Close() error {
	sg.lock.RLock()
	errs := errorx.Errors{}
	for _, svr := range sg.serves {
		err := svr.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	sg.lock.RUnlock()
	sg.wg.Wait()
	return errs.MaybeUnwrap()
}

type ServeWrap struct {
	ServeFn func() error
	CloseFn func() error
}

func (sw *ServeWrap) Serve() error {
	return sw.ServeFn()
}

func (sw *ServeWrap) Close() error {
	return sw.CloseFn()
}

func NewServeWrap(serve func() error, close func() error) ServeItem {
	return &ServeWrap{ServeFn: serve, CloseFn: close}
}
