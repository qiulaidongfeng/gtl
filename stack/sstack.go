// gsstack
package stack

import (
	"sync"
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
	*s.scap = 2
	return s
}

func (s *Sstack) addcap(size uint64) (ncap uint64) {
	ncap = uint64(cap(s.slice))
	for ncap < size {
		s.slice = append(s.slice, 2)
		ncap = uint64(cap(s.slice))
		nslice := make([]int8, ncap, ncap)
		for i := uint64(0); i < (*s.size); i++ {
			nslice[i] = s.slice[i]
		}
		s.slice = nslice
	}
	return
}

func (s *Sstack) Tsaddcap(size uint64) (ncap uint64) {
	s.mutex.Lock()
	ncap = uint64(cap(s.slice))
	for ncap < size {
		s.slice = append(s.slice, 2)
		ncap = uint64(cap(s.slice))
		nslice := make([]int8, ncap, ncap)
		for i := uint64(0); i < (*s.size); i++ {
			nslice[i] = s.slice[i]
		}
		s.slice = nslice
	}
	s.mutex.Unlock()
	return
}

func (s *Sstack) Push(ptr unsafe.Pointer, size uint64) error {
	if (*s.size)+size >= (*s.scap) {
		*s.scap = s.addcap(size)
	}
	// fot i:=0;i<size;i++{
	// 	s.slice[]=
	// }
	return nil
}
