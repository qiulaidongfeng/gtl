// GLMstackpop
package stack

import (
	"sync/atomic"
	"unsafe"
)

func (s *GLMstack) Popptr(ptr unsafe.Pointer, size uint64) error {
	sizei := s.size
	s.size -= size
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	uptr := uintptr(ptr)
	for i := uint64(0); i < size; i++ {
		*(*int8)(unsafe.Pointer(vptr + (uintptr(i)))) = *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
	}
	return nil
}

func (s *GLMstack) TsPopptr(ptr unsafe.Pointer, size uint64) error {
	s.mutex.RLock()
	s.poprecord() //出栈记录
	sizei := atomic.AddUint64(&s.size, ^(size - 1))
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	uptr := uintptr(ptr)
	for i := uint64(0); i < size; i++ { //实际出栈
		*(*int8)(unsafe.Pointer(vptr + (uintptr(i)))) = *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
	}
	s.popok() //结束记录
	s.mutex.RUnlock()
	return nil
}
