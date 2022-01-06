// gsstack
package stack

import (
	"sync"
)

type TypeCode byte

const (
	Int8code       = TypeCode(0)
	Int16code      = TypeCode(1)
	Int32code      = TypeCode(2)
	Int64code      = TypeCode(3)
	Intcode        = TypeCode(4)
	Uint8code      = TypeCode(5)
	Uint16code     = TypeCode(6)
	Uint32code     = TypeCode(7)
	Uint64code     = TypeCode(8)
	Uintcode       = TypeCode(9)
	Bytecode       = TypeCode(10)
	Runecode       = TypeCode(11)
	Stringcode     = TypeCode(12)
	Interfacecode  = TypeCode(13)
	Chancode       = TypeCode(14)
	Mapcode        = TypeCode(15)
	Boolcode       = TypeCode(16)
	Float32code    = TypeCode(17)
	Float64code    = TypeCode(18)
	Uintptrcode    = TypeCode(19)
	Complex64code  = TypeCode(20)
	Complex128code = TypeCode(21)
)

type GenericitySavememoryStack struct {
	slice []int8
	size  *uint64
	mutex sync.RWMutex
}

type Sstack = GenericitySavememoryStack

func NewGsstack() Gsstack {
	s := Gsstack{
		slice: make([]int8, 0, 2),
		size:  new(uint64),
	}
	return s
}

func (s *Gsstack) Pushint8(x int8) {
	s.slice = append(s.slice[:((*s.size)+1)], x)
	*s.size++
	return
}

func (s *Gsstack) Push(x interface{}, Type TypeCode) {
	s.slice = append(s.slice[:((*s.size)+1)], x)
	*s.size++
	return
}
