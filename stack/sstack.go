// gsstack
package stack

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	Int8size       uintptr = (unsafe.Sizeof(int8(1)))
	Int16size      uintptr = (unsafe.Sizeof(int16(1)))
	Int32size      uintptr = (unsafe.Sizeof(int32(1)))
	Int64size      uintptr = (unsafe.Sizeof(int64(1)))
	Intsize        uintptr = (unsafe.Sizeof(int(1)))
	Uint8size      uintptr = (unsafe.Sizeof(uint8(1)))
	Uint16size     uintptr = (unsafe.Sizeof(uint16(1)))
	Uint32size     uintptr = (unsafe.Sizeof(uint32(1)))
	Uint64size     uintptr = (unsafe.Sizeof(uint64(1)))
	Uintsize       uintptr = (unsafe.Sizeof(uint(1)))
	Bytesize       uintptr = (unsafe.Sizeof(byte(1)))
	Runesize       uintptr = (unsafe.Sizeof(rune(1)))
	Boolsize       uintptr = (unsafe.Sizeof(bool(true)))
	Float32size    uintptr = (unsafe.Sizeof(float32(1.0)))
	Float64size    uintptr = (unsafe.Sizeof(float64(2.0)))
	Uintptrsize    uintptr = (unsafe.Sizeof(uintptr(7)))
	Complex64size  uintptr = (unsafe.Sizeof(complex64(6 + 9i)))
	Complex128size uintptr = (unsafe.Sizeof(complex128(8 + 9i)))
	Interfacesize  uintptr = (unsafe.Sizeof(*(new(interface{}))))
)

type GenericitySavememoryStack struct {
	slice []int8
	size  *uint64
	scap  *uint64
	mutex sync.RWMutex
}

type Sstack = GenericitySavememoryStack

func NewGsstack() Sstack {
	s := Sstack{
		slice: make([]int8, 2, 2),
		size:  new(uint64),
		scap:  new(uint64),
	}
	s.scap = 2
	return s
}

func (s *Sstack) addcap() (ncap uint64) {
	s.slice = append(s.slice, 9)
	ncap = cap(s.slice)
	nslice := make([]int8, ncap, ncap)
	for i := 0; i < (ncap - 1); i++ {
		nslice[i] = s.slice[i]
	}
	return
}

func (s *Sstack) Tsaddcap() (ncap uint64) {
	s.mutex.Lock()
	s.slice = append(s.slice, 9)
	ncap = cap(s.slice)
	nslice := make([]int8, ncap, ncap)
	for i := 0; i < (ncap - 1); i++ {
		nslice[i] = s.slice[i]
	}
	s.mutex.Ulock()
	return
}

func (s *Sstack) Pushint8(x int8) error {
	if (*s.scap) == (*s.size) {
		*s.scap = s.addcap()
	}
	s.slice[size] = x
	s.size++
	return nil
}

func (s *Sstack) TsPushint8(x int8) {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int8size)
	ok := atomic.CompareAndSwapUint64(s.scap, nsize-Int8size, ((*s.scap) + 100))
	if ok {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap()
		s.mutex.RLock()
	} else {
		s.mutex.RUnlock()
	}
	s.slice[nsize] = x
	return nil
}

func (s *Sstack) Pushint16(x int16) {
	if (*s.scap) == (*s.size) {
		*s.scap = s.addcap()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	s.size += Int16size
	return
}

func (s *Sstack) Pushint32(x int32) {

}

// func (s *Sstack) Push(x interface{}, Type TypeCode) {
// 	s.slice = append(s.slice[:((*s.size)+1)], x)
// 	*s.size++
// 	return
// }
