package util

import "sync"

func WithLock(lock sync.Locker, f func()) {
	lock.Lock()
	defer lock.Unlock()
	f()
}
