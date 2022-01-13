// lmrwmutex
package gsync

import (
	"errors"
	"runtime"
	"sync/atomic"
)

var (
	Noreadlock  error = errors.New("No read lock")
	Nowritelock error = errors.New("No write lock")
)

type LMrwmutex struct {
	nm int64
}

func (m *LMrwmutex) Lock() {
	for {
		ok := atomic.CompareAndSwapInt64(&m.nm, 0, -1)
		if ok == true {
			break
		}
		runtime.Gosched()
	}
}

func (m *LMrwmutex) Unlock() {
	nm := atomic.LoadInt64(&m.nm)
	if nm >= 0 {
		panic(Nowritelock)
	}
	for {
		ok := atomic.CompareAndSwapInt64(&m.nm, -1, 0)
		if ok == true {
			break
		}
		runtime.Gosched()
	}
}

func (m *LMrwmutex) RLock() {
	for {
		ok := atomic.CompareAndSwapInt64(&m.nm, -1, -1)
		if ok == true {
			runtime.Gosched()
			continue
		}
		nm := atomic.AddInt64(&m.nm, 1)
		if nm == 0 {
			break
		}
	}
}

func (m *LMrwmutex) RUnlock() {
	ok := atomic.CompareAndSwapInt64(&m.nm, -1, -1)
	if ok == true {
		panic(Noreadlock)
	}
	ok = atomic.CompareAndSwapInt64(&m.nm, 0, 0)
	if ok == true {
		panic(Noreadlock)
	}
	for {
		atomic.AddInt64(&m.nm, -1)
		break
	}
}
