package sys

import (
	"syscall"
)

//syscall.Handle别名
type Handle = syscall.Handle

//BUG(Dll): 在windows/amd64环境下的go test测试中使用DLL可能导致不能recover的异常

//动态链接库结构体
type DLL struct {
	name string
	dll  Handle
}

//打开一个动态链接库
func NewDLL(name string) (d *DLL, err error) {
	d = new(DLL)
	d.name = name
	d.dll, err = syscall.LoadLibrary(name)
	if err != nil {
		return nil, err
	}
	return
}

//寻找一个动态链接库中导出的过程
func (d *DLL) FindProc(name string) (proc uintptr, err error) {
	return syscall.GetProcAddress(d.dll, name)
}

//关闭一个动态链接库
func (d *DLL) Release() (err error) {
	syscall.FreeLibrary(d.dll)
	return
}

//关闭一个动态链接库
func (d *DLL) Close() (err error) {
	syscall.FreeLibrary(d.dll)
	return
}

func (d *DLL) String() string {
	return d.name + ":DLL"
}
