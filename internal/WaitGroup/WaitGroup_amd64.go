package WaitGroup

import (
	"sync"
	"sync/atomic"
)

type WaitGroup struct {
	wg   sync.WaitGroup
	n    int64
	oldn int64
}

func (wg *WaitGroup) Add(delta int) {
	atomic.AddInt64(&wg.n, int64(delta))
	atomic.AddInt64(&wg.oldn, int64(delta))
	wg.wg.Add(int(delta))
}

func (wg *WaitGroup) Done() {
	n := atomic.AddInt64(&wg.n, -1)
	if n == 0 {
		wg.wg.Add(-int(atomic.LoadInt64(&wg.oldn)))
	}
}

func (wg *WaitGroup) Wait() {
	wg.wg.Wait()
}

func (wg *WaitGroup) Init() {
	wg.n = 0
	wg.oldn = 0
}
