// gsstack
package stack

import (
	"sync"
	"sync/atomic"
	"unsafe"
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

type GenericitySavememoryStack struct {
	slice  []int8
	size   *uint64
	scap   *uint64
	bmutex bool
	mutex  sync.RWMutex
}

type Sstack = GenericitySavememoryStack

func NewGsstack() Sstack {
	s := Sstack{
		slice:  make([]int8, 2, 2),
		size:   new(uint64),
		scap:   new(uint64),
		bmutex: false,
	}
	*s.scap = 2
	return s
}

func (s *Sstack) addcap() (ncap uint64) {
	s.slice = append(s.slice, 2)
	ncap = uint64(cap(s.slice))
	nslice := make([]int8, ncap, ncap)
	for i := uint64(0); i < (ncap - 1); i++ {
		nslice[i] = s.slice[i]
	}
	s.slice = nslice
	return
}

func (s *Sstack) Tsaddcap() (ncap uint64) {
	s.mutex.Lock()
	s.slice = append(s.slice, 9)
	ncap = uint64(cap(s.slice))
	nslice := make([]int8, ncap, ncap)
	for i := uint64(0); i < (ncap - 1); i++ {
		nslice[i] = s.slice[i]
	}
	s.mutex.Unlock()
	return
}

func (s *Sstack) Pushint8(x int8) error {
	if (*s.size)+Int8size >= (*s.scap) {
		*s.scap = s.addcap()
	}
	s.slice[*s.size] = x
	*s.size++
	return nil
}

func (s *Sstack) TsPushint8(x int8) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int8size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap()
		s.mutex.RLock()
	} else {
		s.mutex.RUnlock()
	}
	s.slice[nsize] = x
	return nil
}

func (s *Sstack) Pushint16(x int16) error {
	if (*s.size)+Int16size >= (*s.scap) {
		*s.scap = s.addcap()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	*s.size += Int16size
	return nil
}

func (s *Sstack) Pushint32(x int32) {

}

// func (s *Sstack) Push(x interface{}, Type TypeCode) {
// 	s.slice = append(s.slice[:((*s.size)+1)], x)
// 	*s.size++
// 	return
// }
