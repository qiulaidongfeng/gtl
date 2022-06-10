package Copys

import (
	"unsafe"
)

func Copy_noasm(dest, src unsafe.Pointer, n uint) {
	for i := uint(0); i < n; i++ {
		*(*int8)((unsafe.Pointer(((uintptr)(dest)) + uintptr(i)))) = *(*int8)((unsafe.Pointer(((uintptr)(src)) + uintptr(i))))
	}
	return
}
