// gsstack
package stack

import (
	"sync"
	"unsafe"
)

const (
	Int8code       uintptr = (unsafe.Sizeof(int8(1)))
	Int16code      uintptr = (unsafe.Sizeof(int16(1)))
	Int32code      uintptr = (unsafe.Sizeof(int32(1)))
	Int64code      uintptr = (unsafe.Sizeof(int64(1)))
	Intcode        uintptr = (unsafe.Sizeof(int(1)))
	Uint8code      uintptr = (unsafe.Sizeof(uint8(1)))
	Uint16code     uintptr = (unsafe.Sizeof(uint16(1)))
	Uint32code     uintptr = (unsafe.Sizeof(uint32(1)))
	Uint64code     uintptr = (unsafe.Sizeof(uint64(1)))
	Uintcode       uintptr = (unsafe.Sizeof(uint(1)))
	Bytecode       uintptr = (unsafe.Sizeof(byte(1)))
	Runecode       uintptr = (unsafe.Sizeof(rune(1)))
	Boolcode       uintptr = (unsafe.Sizeof(bool(true)))
	Float32code    uintptr = (unsafe.Sizeof(float32(1.0)))
	Float64code    uintptr = (unsafe.Sizeof(float64(2.0)))
	Uintptrcode    uintptr = (unsafe.Sizeof(uintptr(7)))
	Complex64code  uintptr = (unsafe.Sizeof(complex64(6 + 9i)))
	Complex128code uintptr = (unsafe.Sizeof(complex128(8 + 9i)))
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
		slice: make([]int8, 0, 2),
		size:  new(uint64),
		scap:  new(uint64),
	}
	s.scap = 2
	return s
}

func (s *Sstack) Pushint8(x int8) {
	s.slice = append(s.slice[:((*s.size)+1)], x)
	*s.size++
	return
}

func (s *Sstack) Pushint16(x int16) {
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl) + uintptr(1))
	sp2 := (*int16)(sp)
	*sp2 = x
	return
}

func (s *Sstack) Pushint32(x int32) {

}

// func (s *Sstack) Push(x interface{}, Type TypeCode) {
// 	s.slice = append(s.slice[:((*s.size)+1)], x)
// 	*s.size++
// 	return
// }
