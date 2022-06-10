package Copys

import (
	"unsafe"
)

//go:nosplit
func div(a, b uint64) (c uint64, d uint64)

//go:nosplit
func Copy_MOVSQ(dest, src unsafe.Pointer, n uint)

//go:nosplit
func Copy_SSE_Movups(dest, src unsafe.Pointer, n uint)

//go:uintptrescapes
func mtob(addr, length uintptr) []byte {
	var a []byte
	p := (*[3]uintptr)(unsafe.Pointer(&a))
	p[0] = addr
	p[1] = length
	p[2] = length
	return a
}
