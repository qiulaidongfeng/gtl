// gsstack
package stack

import (
	"sync"
	"unsafe"
)

const (
	Int8code       int = int(unsafe.Sizeof(int8(1)))
	Int16code      int = int(unsafe.Sizeof(int16(1)))
	Int32code      int = int(unsafe.Sizeof(int32(1)))
	Int64code      int = int(unsafe.Sizeof(int64(1)))
	Intcode        int = int(unsafe.Sizeof(int(1)))
	Uint8code      int = int(unsafe.Sizeof(uint8(1)))
	Uint16code     int = int(unsafe.Sizeof(uint16(1)))
	Uint32code     int = int(unsafe.Sizeof(uint32(1)))
	Uint64code     int = int(unsafe.Sizeof(uint64(1)))
	Uintcode       int = int(unsafe.Sizeof(uint(1)))
	Bytecode       int = int(unsafe.Sizeof(byte(1)))
	Runecode       int = int(unsafe.Sizeof(rune(1)))
	Boolcode       int = int(unsafe.Sizeof(bool(true)))
	Float32code    int = int(unsafe.Sizeof(float32(1.0)))
	Float64code    int = int(unsafe.Sizeof(float64(2.0)))
	Uintptrcode    int = int(unsafe.Sizeof(uintptr(7)))
	Complex64code  int = int(unsafe.Sizeof(complex64(6 + 9i)))
	Complex128code int = int(unsafe.Sizeof(complex128(8 + 9i)))
)

type GenericitySavememoryStack struct {
	slice []int8
	size  *uint64
	slen  *uint64
	mutex sync.RWMutex
}

type Sstack = GenericitySavememoryStack

func NewGsstack() Sstack {
	s := Sstack{
		slice: make([]int8, 0, 2),
		size:  new(uint64),
		slen:  new(uint64),
	}
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
