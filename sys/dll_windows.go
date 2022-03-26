package sys

import (
	"syscall"

	"golang.org/x/sys/windows"
)

type Errno = syscall.Errno

//动态链接库结构体
type DLL struct {
	name string
	dll  *windows.DLL
}

//打开一个动态链接库
func NewDLL(name string) (d *DLL, err error) {
	d.name = name
	d.dll, err = windows.LoadDLL(name)
	if err != nil {
		return nil, err
	}
	return
}

//寻找一个动态链接库中导出的过程
func (d *DLL) FindProc(name string) (proc uintptr, err error) {
	var Proc *windows.Proc
	Proc, err = d.dll.FindProc(name)
	if err != nil {
		return 0, err
	}
	proc = Proc.Addr()
	return
}

//关闭一个动态链接库
func (d *DLL) Release() (err error) {
	err = d.dll.Release()
	return
}

//调用指定名称的动态链接库中导出的过程
func (d *DLL) Call(name string, a ...uintptr) (r1, r2 uintptr, err error) {
	var procaddr uintptr
	procaddr, err = d.FindProc(name)
	if err != nil {
		return 0, 0, err
	}
	return syscall.SyscallN(procaddr, a...)
}

//调用已知地址的动态链接库中导出的过程
func Call(procaddr uintptr, a ...uintptr) (r1, r2 uintptr, err Errno) {
	return syscall.SyscallN(procaddr, a...)
}

func (d *DLL) String() string {
	return d.name + ":DLL"
}
