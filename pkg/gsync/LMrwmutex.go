package gsync

import (
	"errors"
	"runtime"
	"sync/atomic"
)

// 读写锁错误
var (
	//读锁没有被持有
	Noreadlock error = errors.New("No read lock")
	//写锁没有被持有
	Nowritelock error = errors.New("No write lock")
)

const (
	nolock    int64 = 0
	writelock int64 = -1 << 63
)

type nocopy struct{}

func (_ nocopy) Lock()   {}
func (_ nocopy) Unlock() {}

// 使用内存少的读写锁，不能复制
type LowMemory_rwMutex struct {
	_     nocopy
	state int64
}

// 使用内存少的读写锁，不能复制
type LMrwmutex = LowMemory_rwMutex

// 获取写锁
func (m *LMrwmutex) Lock() {
	for {
		if atomic.CompareAndSwapInt64(&m.state, nolock, writelock) {
			break
		}
		runtime.Gosched()
	}
}

// 释放写锁
func (m *LMrwmutex) Unlock() {
	if !atomic.CompareAndSwapInt64(&m.state, writelock, nolock) {
		panic(Nowritelock)
	}
}

// 获取读锁
func (m *LMrwmutex) RLock() {
	for {
		if atomic.AddInt64(&m.state, 1) > 0 { //如果现在是读锁
			return
		}
		atomic.AddInt64(&m.state, -1)
		runtime.Gosched()
	}
}

// 释放读锁
func (m *LMrwmutex) RUnlock() {
	if atomic.AddInt64(&m.state, -1) < 0 {
		panic(Noreadlock)
	}
}
