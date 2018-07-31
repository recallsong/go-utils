package servegrp

import (
	"sync"
)

type WaitClose struct {
	sync.WaitGroup
	CloseCh chan struct{}
}

func NewWaitClose() *WaitClose {
	return &WaitClose{
		CloseCh: make(chan struct{}),
	}
}

func (wc *WaitClose) Close() error {
	close(wc.CloseCh)
	wc.Wait()
	return nil
}
