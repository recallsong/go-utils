package dic

import "sync"

// SyncDicPair 带锁的Dic类型
type SyncDicPair struct {
	Dic
	sync.RWMutex
}
