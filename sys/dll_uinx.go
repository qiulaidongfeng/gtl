//go:build !windows
// +build !windows

package sys

import (
	"errors"
	"unsafe"
)

//动态链接库结构体
type DLL struct {
	name string
	addr unsafe.Pointer
}

//打开一个动态链接库
func NewDLL(name string) (d *DLL, err error) {
	return NewDLLAll(name, RTLD_NOW)
}

//打开一个动态链接库,以指定的mode
func NewDLLAll(name string, mode int) (d *DLL, err error) {
	d.name = name
	d.addr = Dlopen(name, mode)
	if d.addr == nil {
		return nil, errors.New(Dlerror())
	}
	return
}

//关闭一个动态链接库
func (d *DLL) Close() (errint int, err error) {
	ret := Dlclose(d.addr)
	if ret != 0 {
		return ret, errors.New(Dlerror())
	}
	return ret, nil
}

//寻找一个动态链接库中导出的过程
func (d *DLL) FindProc(name string) (proc uintptr, err error) {
	var ptr unsafe.Pointer
	ptr = Dlsym(d.addr, name)
	if ptr == nil {
		return 0, errors.New(Dlerror())
	}
	return uintptr(ptr), nil
}

func (d *DLL) String() string {
	return d.name + ".so"
}
