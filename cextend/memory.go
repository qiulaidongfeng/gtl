// memory
package cextend

/*
	#cgo CFLAGS: -g -O3 -march=corei7
	#include "cextend.h"
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

func Realloc(ptr unsafe.Pointer, size uint) (nptr unsafe.Pointer) {
	nptr = C.realloc(ptr, C.size_t(size))
	return
}

func Malloc(size uint) (ptr unsafe.Pointer) {
	ptr = C.malloc(C.size_t(size))
	return
}

func Calloc(nitems uint, size uint) (ptr unsafe.Pointer) {
	ptr = C.calloc(C.size_t(nitems), C.size_t(size))
	return
}

func Free(ptr unsafe.Pointer) {
	C.free(ptr)
	return
}

func Memcpy(dest, src unsafe.Pointer, n uint) {
	gcdest := cgo.NewHandle(dest)
	cdest := gcdest.Value().(unsafe.Pointer)
	gcsrc := cgo.NewHandle(src)
	csrc := gcsrc.Value().(unsafe.Pointer)
	defer gcdest.Delete()
	defer gcsrc.Delete()
	C.memcpy(cdest, csrc, C.size_t(n))
	return
}
