package gsync

import (
	"errors"
	"runtime"
	"sync/atomic"
)

//读写锁错误
var (
	//读锁没有被持有
	Noreadlock error = errors.New("No read lock")
	//写锁没有被持有
	Nowritelock error = errors.New("No write lock")
)

const (
	nolock    = 0
	writelock = -1
)

type nocopy struct{}

//使用内存少的读写锁，不能复制
type LowMemory_rwMutex struct {
	_  nocopy
	nm int64
}

//使用内存少的读写锁，不能复制
type LMrwmutex = LowMemory_rwMutex

//获取写锁
func (m *LMrwmutex) Lock() {
	for {
		ok := atomic.CompareAndSwapInt64(&m.nm, nolock, writelock)
		if ok == true {
			break
		}
		runtime.Gosched()
	}
}

//释放写锁
func (m *LMrwmutex) Unlock() {
	nm := atomic.LoadInt64(&m.nm)
	if nm >= 0 {
		panic(Nowritelock)
	}
	atomic.StoreInt64(&m.nm, nolock)
}

//获取读锁
func (m *LMrwmutex) RLock() {
	for {
		ok := atomic.CompareAndSwapInt64(&m.nm, writelock, writelock)
		if ok == true {
			runtime.Gosched()
			continue
		}
		atomic.AddInt64(&m.nm, 1)
		break
	}
}

//释放读锁
func (m *LMrwmutex) RUnlock() {
	nm := atomic.LoadInt64(&m.nm)
	if nm == -1 {
		panic(Noreadlock)
	} else if nm == 0 {
		panic(Noreadlock)
	}
	atomic.AddInt64(&m.nm, -1)
}
