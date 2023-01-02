//go:build !windows
// +build !windows

package sys

import (
	"syscall"
	"unsafe"
)

// 指定虚拟内存操作的结构体
type Virtual_memory struct {
	addr   uintptr
	length uintptr
}

// 指定虚拟内存操作的结构体别名
type VM = Virtual_memory

// 创建指定虚拟内存操作的结构体
func NewVM(addr, length uintptr) VM {
	return VM{
		addr:   addr,
		length: length,
	}
}

func mtob(addr, length uintptr) []byte {
	var a []byte
	p := (*[3]uintptr)(unsafe.Pointer(&a))
	p[0] = addr
	p[1] = length
	p[2] = length
	return a
}

// 锁定指定虚拟内存操作，有错panic
func (v VM) Lock() {
	pslice := mtob(v.addr, v.length)
	err := syscall.Mlock(pslice)
	if err != nil {
		panic(err)
	}
}

// 锁定指定虚拟内存操作，有错返回error
func (v VM) Lockerr() error {
	pslice := mtob(v.addr, v.length)
	err := syscall.Mlock(pslice)
	return err
}

// 解锁指定虚拟内存操作，有错panic
func (v VM) Unlock() {
	pslice := mtob(v.addr, v.length)
	err := syscall.Munlock(pslice)
	if err != nil {
		panic(err)
	}
}

// 解锁指定虚拟内存操作，有错error
func (v VM) Unlockerr() error {
	pslice := mtob(v.addr, v.length)
	err := syscall.Munlock(pslice)
	return err
}
