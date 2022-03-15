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

//C语言<stdlib.h>中提供的realloc函数的go语言API
func Realloc(ptr unsafe.Pointer, size uint) (nptr unsafe.Pointer) {
	nptr = C.realloc(ptr, C.size_t(size))
	return
}

//C语言<stdlib.h>中提供的malloc函数的go语言API
func Malloc(size uint) (ptr unsafe.Pointer) {
	ptr = C.malloc(C.size_t(size))
	return
}

//C语言<stdlib.h>中提供的calloc函数的go语言API
func Calloc(nitems uint, size uint) (ptr unsafe.Pointer) {
	ptr = C.calloc(C.size_t(nitems), C.size_t(size))
	return
}

//C语言<stdlib.h>中提供的free函数的go语言API
func Free(ptr unsafe.Pointer) {
	C.free(ptr)
	return
}

//C语言<string.h>中提供的memcpy函数的go语言API
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
