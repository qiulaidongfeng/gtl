// memory
package cextend

/*
	#cgo CFLAGS: -g -O3 -march=corei7 -fPIC -I./C
	#include <cextend.h>
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

func Malloc(size uint) (ptr unsafe.Pointer) {
	ptr = C.Malloc(C.ulong(size))
	return
}

func Free(ptr unsafe.Pointer) {
	C.Free(ptr)
	return
}

func Memcpy(dest, src unsafe.Pointer, n uint) {
	gcdest := cgo.NewHandle(dest)
	cdest := gcdest.Value().(unsafe.Pointer)
	gcsrc := cgo.NewHandle(src)
	csrc := gcsrc.Value().(unsafe.Pointer)
	C.Memcpy(C.uintptr_t(uintptr(cdest)), C.uintptr_t(uintptr(csrc)), C.ulong(n))
	gcdest.Delete()
	gcsrc.Delete()
	return
}
