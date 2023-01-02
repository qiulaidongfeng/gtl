//go:build !amd64
// +build !amd64

package Copys

import (
	"unsafe"
)

func Copy(dest, src unsafe.Pointer, n uint) {
	n8 := n / 8
	n8n := n8 * 8
	var n1 uint
	if n8n != n {
		n1 = n8 - n8n
	}
	off := 0
	for i := uint(0); i < n8; i++ {
		*(*int64)((unsafe.Pointer(((uintptr)(dest)) + uintptr(off)))) = *(*int64)((unsafe.Pointer(((uintptr)(src)) + uintptr(off))))
		off += 8
	}
	for i := uint(0); i < n1; i++ {
		*(*int8)((unsafe.Pointer(((uintptr)(dest)) + uintptr(i)))) = *(*int8)((unsafe.Pointer(((uintptr)(src)) + uintptr(i))))
	}
	return
}
