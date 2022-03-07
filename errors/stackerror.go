// stckerror
package errors

import (
	"fmt"
	"runtime"
)

var (
	//单goroutine栈踪迹信息大小
	StackSize = 1024
	//所有goroutine栈踪迹信息大小
	AllStackSize = runtime.NumGoroutine() * StackSize
)

//带有栈踪迹信息的错误
type WrapStackError struct {
	err   error
	stack []byte
}

//创建带有栈踪迹信息的错误
func NewWrapStackError(err string, all bool) PlusWrapError {
	errorerr := Errorstring(err)
	return &WrapStackError{err: &errorerr, stack: Stack(all)}
}

//包装当前栈信息(
func WrapFuncStackError(err *WrapStackError) {
	err.stack = append(callStack(), err.stack...)
}

func (err *WrapStackError) Error() string {
	return err.err.Error() + "\n" + string(err.stack)
}

func (err *WrapStackError) Unwrap() error {
	return err.err
}

func (err *WrapStackError) ErrorNoStack() string {
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

func callStack() []byte {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return nil
	}
	call := runtime.FuncForPC(pc)
	file, line := call.FileLine(pc)
	return []byte(fmt.Sprint(call.Name(), "()", "\n\t", file, line))
}
