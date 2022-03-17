//go:build linux
// +build linux

package main

import (
	"unsafe"
)

/*
	#include <dlfcn.h>
*/
import "C"

const (
	/*
		重定位是在依赖于实现的时间执行的,范围从 dlopen() 调用时间到对给定符号的第一次引用发生
		指定 RTLD_LAZY 应该会提高支持动态符号绑定的实现的性能，因为进程可能不会引用任何给定对象中的所有函数
		而且，对于支持正常流程执行的动态符号解析的系统，此行为模仿流程执行的正常处理。
	*/
	RTLD_LAZY = C.RTLD_LAZY
	/*
		加载对象时执行重定位。
		首次加载对象时会执行所有必要的重定位。
		如果对从未引用的函数执行重定位，这可能会浪费一些处理。
		对于需要在加载对象后立即知道执行期间引用的所有符号都可用的应用程序，此行为可能很有用。
	*/
	RTLD_NOW = C.RTLD_NOW
	/*
		所有符号都可用于其他模块的重定位处理。
		对象的符号可用于任何其他对象的重定位处理。
		此外，使用 dlopen ( 0, mode ) 和关联的 dlsym()进行符号查找允许搜索使用此模式 加载的对象 。
	*/
	RTLD_GLOBAL = C.RTLD_GLOBAL
	/*
		并非所有符号都可用于其他模块的重定位处理。
		对象的符号不可用于任何其他对象的重定位处理。
	*/
	RTLD_LOCAL = C.RTLD_LOCAL
)

//C语言#include <dlfcn.h>提供的dlopen函数的go语言API
func Dlopen(file string, mode int) (ptr unsafe.Pointer) {
	return C.dlopen(C.CString(file), C.int(mode))
}

//C语言#include <dlfcn.h>提供的dlsym函数的go语言API
func Dlsym(handle unsafe.Pointer, name string) (ptr unsafe.Pointer) {
	return C.dlsym(handle, C.CString(name))
}

//C语言#include <dlfcn.h>提供的dlclose函数的go语言API
func Dlclose(handle unsafe.Pointer) int {
	ret := C.dlclose(handle)
	return int(ret)
}

//C语言#include <dlfcn.h>提供的dlerror函数的go语言API
func Dlerror() string {
	ret := C.dlerror()
	return string(ret)
}
