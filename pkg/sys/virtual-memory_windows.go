package sys

import (
	"syscall"
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

// 锁定指定虚拟内存操作，有错panic
func (v VM) Lock() {
	err := syscall.VirtualLock(v.addr, v.length)
	if err != nil {
		panic(err)
	}
}

// 锁定指定虚拟内存操作，有错返回error
func (v VM) Lockerr() error {
	err := syscall.VirtualLock(v.addr, v.length)
	return err
}

// 解锁指定虚拟内存操作，有错panic
func (v VM) Unlock() {
	err := syscall.VirtualUnlock(v.addr, v.length)
	if err != nil {
		panic(err)
	}
}

// 解锁指定虚拟内存操作，有错error
func (v VM) Unlockerr() error {
	err := syscall.VirtualUnlock(v.addr, v.length)
	return err
}
