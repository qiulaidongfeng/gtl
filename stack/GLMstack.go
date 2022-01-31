// gsstack
package stack

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	pushpp int64 = -1
	poppp  int64 = 1
)

const (
	Poptime  = int64(70)
	Pushtime = int64(140)
	Waittime = int64(200)
)

const (
	Int8size       uint64 = uint64((unsafe.Sizeof(int8(1))))
	Int16size      uint64 = uint64((unsafe.Sizeof(int16(1))))
	Int32size      uint64 = uint64((unsafe.Sizeof(int32(1))))
	Int64size      uint64 = uint64((unsafe.Sizeof(int64(1))))
	Intsize        uint64 = uint64((unsafe.Sizeof(int(1))))
	Uint8size      uint64 = uint64((unsafe.Sizeof(uint8(1))))
	Uint16size     uint64 = uint64((unsafe.Sizeof(uint16(1))))
	Uint32size     uint64 = uint64((unsafe.Sizeof(uint32(1))))
	Uint64size     uint64 = uint64((unsafe.Sizeof(uint64(1))))
	Uintsize       uint64 = uint64((unsafe.Sizeof(uint(1))))
	Bytesize       uint64 = uint64((unsafe.Sizeof(byte(1))))
	Runesize       uint64 = uint64((unsafe.Sizeof(rune(1))))
	Boolsize       uint64 = uint64((unsafe.Sizeof(bool(true))))
	Float32size    uint64 = uint64((unsafe.Sizeof(float32(1.0))))
	Float64size    uint64 = uint64((unsafe.Sizeof(float64(2.0))))
	Uintptrsize    uint64 = uint64((unsafe.Sizeof(uintptr(7))))
	Complex64size  uint64 = uint64((unsafe.Sizeof(complex64(6 + 9i))))
	Complex128size uint64 = uint64((unsafe.Sizeof(complex128(8 + 9i))))
	Interfacesize  uint64 = uint64((unsafe.Sizeof(*(new(interface{})))))
)

type GenericLowMemoryStack struct {
	slice    []int8
	size     uint64
	scap     uint64
	pushn    int64
	popn     int64
	pp       int64
	pushtime int64
	poptime  int64
	waittime int64
	mutex    sync.RWMutex
}

type GLMstack = GenericLowMemoryStack

func NewGLMstack() *GLMstack {
	s := &GLMstack{
		slice:    make([]int8, 2, 2),
		size:     2,
		scap:     2,
		pushn:    0,
		popn:     0,
		pp:       0,
		pushtime: Pushtime,
		poptime:  Poptime,
		waittime: Waittime,
	}
	return s
}

func (s *GLMstack) addcap(size uint64) (ncap uint64) {
	ncap = uint64(cap(s.slice))
	for ncap <= size {
		if ncap < 1024 {
			ncap += ncap
		} else {
			ncap += ncap / 4
		}
	}
	nslice := make([]int8, ncap, ncap)
	// nptr := uintptr(unsafe.Pointer(&nslice[0]))
	// uptr := uintptr(unsafe.Pointer(&s.slice[0]))
	for i := uint64(0); i < s.scap; i++ {
		nslice[i] = s.slice[i]
	}
	s.slice = nslice
	return
}

func (s *GLMstack) pushsafetycheck(size uint64) {
	if s.size+size >= s.scap {
		s.scap = s.addcap(s.size + size)
	}
}

func (s *GLMstack) tspushsafetycheck(size uint64) (sizeold uint64) {
	nsize := atomic.AddUint64(&s.size, size)
	if nsize >= s.scap {
		s.mutex.RUnlock()
		s.scap = s.addcap(nsize)
		s.mutex.RLock()
	}
	sizeold = nsize - size
	return
}

func (s *GLMstack) tsaddcap(size uint64) (ncap uint64) {
	s.mutex.Lock()
	ncap = uint64(cap(s.slice))
	for ncap <= size {
		if ncap < 1024 {
			ncap += ncap
		} else {
			ncap += ncap / 4
		}
	}
	nslice := make([]int8, ncap, ncap)
	for i := uint64(0); i < s.scap; i++ {
		nslice[i] = s.slice[i]
	}
	s.slice = nslice
	s.mutex.Unlock()
	return
}

func (s *GLMstack) Size() uint64 {
	return s.size
}

func (s *GLMstack) Tssize() uint64 {
	return atomic.LoadUint64(&s.size)
}

func (s *GLMstack) Clear() error {
	s.scap = 2
	s.size = 2
	s.slice = make([]int8, 2, 2)
	return nil
}

func (s *GLMstack) Tsclear() error {
	s.mutex.Lock()
	s.scap = 2
	s.size = 2
	s.slice = make([]int8, 2, 2)
	s.mutex.Unlock()
	return nil
}

func (s *GLMstack) Popptr(ptr *unsafe.Pointer, size uint64) error {
	sizei := s.size
	s.size -= size
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	uptr := uintptr(*ptr)
	for i := uint64(0); i < size; i++ {
		*(*int8)(unsafe.Pointer(vptr + (uintptr(i)))) = *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
	}
	return nil
}

func (s *GLMstack) TsPopptr(ptr *unsafe.Pointer, size uint64) error {
	s.mutex.RLock()
	s.poprecord() //出栈记录
	sizei := atomic.AddUint64(&s.size, ^(size - 1))
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	uptr := uintptr(*ptr)
	for i := uint64(0); i < size; i++ { //实际出栈
		*(*int8)(unsafe.Pointer(vptr + (uintptr(i)))) = *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
	}
	s.popok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) TsLoadpp() int64 {
	return atomic.LoadInt64(&s.pp)
}

func (s *GLMstack) TsLoadpopn() int64 {
	return atomic.LoadInt64(&s.popn)
}

func (s *GLMstack) TsLoadpushn() int64 {
	return atomic.LoadInt64(&s.pushn)
}
