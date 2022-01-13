// lmrwmutex
package gsync

import (
	"runtime"
	"sync/atomic"
)

type LMrwmutex struct {
	nm int64
}

func (m *LMrwmutex) Lock() {
	ok := false
	for ok == false {
		ok = atomic.CompareAndSwapInt64(&m.nm, 0, -1)
		runtime.Gosched()
	}
}

func (m *LMrwmutex) Unlock() {
	ok := false
	for ok == false {
		ok = atomic.CompareAndSwapInt64(&m.nm, -1, 0)
		runtime.Gosched()
	}
}
