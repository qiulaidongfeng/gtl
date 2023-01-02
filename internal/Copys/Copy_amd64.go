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
