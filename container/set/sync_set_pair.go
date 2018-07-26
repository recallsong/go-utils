package set

import "sync"

// SyncSetPair 带锁的Set类型
type SyncSetPair struct {
	Set
	sync.RWMutex
}
