// GLMstackpp
package stack

import (
	"sync/atomic"
	"time"
)

//记录为入栈时操作
func (s *GLMstack) pushrecord() {
	atomic.AddInt64(&s.pushn, 1)
	for {
		rw := atomic.LoadInt64(&s.pp)
		if rw == pushpp { //正在入栈操作
			return
		} else if rw == 0 {
			bol := atomic.CompareAndSwapInt64(&s.pp, 0, pushpp)
			if bol == true { //无操作
				return
			} else { //正在出栈操作
				time.Sleep(time.Duration(s.poptime * s.popn))
			}
		} else { //正在出栈操作
			time.Sleep(time.Duration(s.poptime * s.popn))
		}
	}
}

//入栈已完成
func (s *GLMstack) pushok() {
	n := atomic.AddInt64(&s.pushn, -1)
	if n == 0 {
		atomic.StoreInt64(&s.pp, 0)
	}
}

//出栈已完成
func (s *GLMstack) popok() {
	n := atomic.AddInt64(&s.popn, -1)
	if n == 0 {
		atomic.StoreInt64(&s.pp, 0)
	}
}

//记录为出栈时操作
func (s *GLMstack) poprecord() {
	atomic.AddInt64(&s.popn, 1)
	for {
		rw := atomic.LoadInt64(&s.pp)
		if rw == poppp { //正在出栈操作
			return
		} else if rw == 0 {
			bol := atomic.CompareAndSwapInt64(&s.pp, 0, poppp)
			if bol == true { //无操作
				return
			} else {
				time.Sleep(time.Duration(s.pushtime * s.pushn))
			}
		} else {
			time.Sleep(time.Duration(s.pushtime * s.pushn))
		}
	}
}
