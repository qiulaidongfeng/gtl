// memory
package cextend

/*
	#cgo CFLAGS: -g -O3 -march=corei7-avx -static -s -fPIC -fexpensive-optimizations -I./
	#cgo LDFLAGS : -g -O3 -s -march=corei7-avx -static -fPIC -fexpensive-optimizations -L${SRCDIR}/lib -lgocextend_linux64.a
	#include <cextend.h>
*/

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
	C.Memcpy(cdest, csrc, C.ulong(n))
	gcdest.Delete()
	gcsrc.Delete()
	return
}
