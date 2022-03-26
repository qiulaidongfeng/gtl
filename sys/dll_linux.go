package sys

import (
	"errors"
	"strconv"
	"unsafe"

	"github.com/qiulaidongfeng/gtl/cextend"
)

const (
	/*
		重定位是在依赖于实现的时间执行的,范围从 dlopen() 调用时间到对给定符号的第一次引用发生
		指定 RTLD_LAZY 应该会提高支持动态符号绑定的实现的性能，因为进程可能不会引用任何给定对象中的所有函数
		而且，对于支持正常流程执行的动态符号解析的系统，此行为模仿流程执行的正常处理。
	*/
	RTLD_LAZY = cextend.RTLD_LAZY
	/*
		加载对象时执行重定位。
		首次加载对象时会执行所有必要的重定位。
		如果对从未引用的函数执行重定位，这可能会浪费一些处理。
		对于需要在加载对象后立即知道执行期间引用的所有符号都可用的应用程序，此行为可能很有用。
	*/
	RTLD_NOW = cextend.RTLD_NOW
	/*
		所有符号都可用于其他模块的重定位处理。
		对象的符号可用于任何其他对象的重定位处理。
		此外，使用 dlopen ( 0, mode ) 和关联的 dlsym()进行符号查找允许搜索使用此模式 加载的对象 。
	*/
	RTLD_GLOBAL = cextend.RTLD_GLOBAL
	/*
		并非所有符号都可用于其他模块的重定位处理。
		对象的符号不可用于任何其他对象的重定位处理。
	*/
	RTLD_LOCAL = cextend.RTLD_LOCAL
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
	d.addr = cextend.Dlopen(name, mode)
	if d.addr == nil {
		return nil, errors.New(cextend.Dlerror())
	}
	return
}

//关闭一个动态链接库
func (d *DLL) Release() (err int) {
	ret := cextend.Dlclose(d.addr)
	return ret
}

//寻找一个动态链接库中导出的过程
func (d *DLL) FindProc(name string) (proc uintptr, err error) {
	var ptr unsafe.Pointer
	ptr = cextend.Dlsym(d.addr, name)
	if ptr == nil {
		return 0, errors.New(cextend.Dlerror())
	}
	return uintptr(ptr), nil
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
	return d.name + ":so"
}
