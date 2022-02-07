// stckerror
package errors

import (
	"runtime"
)

var (
	//单goroutine栈踪迹信息大小
	StackSize = 1024
	//所有goroutine栈踪迹信息大小
	AllStackSize = runtime.NumGoroutine() * StackSize
)

//带有栈踪迹信息的错误
type StackError struct {
	err   error
	stack string
}

//创建带有栈踪迹信息的错误
func NewStackError(err string, all bool) Error {
	return StackError{err: Errorstring(err), stack: string(Stack(all))}
}

//包装栈踪迹信息进错误
func WrapStackError(err error, all bool) Error {
	return StackError{err: err, stack: string(Stack(all))}
}

func (err StackError) Error() string {
	return err.err.Error() + "\n" + err.stack
}

func (err StackError) Unwrap() error {
	return err.err
}

func (err StackError) ErrorNoStack() string {
	return err.err.Error()
}

//返回栈踪迹信息
func Stack(all bool) []byte {
	if all == false {
		buf := make([]byte, StackSize)
		for {
			n := runtime.Stack(buf, all)
			if n < len(buf) {
				return buf[:n]
			}
			buf = make([]byte, 2*len(buf))
		}
	} else {
		buf := make([]byte, AllStackSize)
		for {
			n := runtime.Stack(buf, all)
			if n < len(buf) {
				return buf[:n]
			}
			buflen := len(buf)
			buf = make([]byte, (buflen + buflen/4))
		}
	}
}
