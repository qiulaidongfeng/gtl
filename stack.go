package stack

import (
	"errors"
)

var (
	StackEmpty           error = errors.New("Stack,Empty")
	StackSizeExceeded    error = errors.New("Stack size exceeded")
	StackPushFail        error = errors.New("Stack push fail")
	StackPopFail         error = errors.New("Stack pop fail")
	StackClearFail       error = errors.New("Stack clear fail")
	StackContentShortage error = NewStackError(StackPopFail, "Stack content shortage")
)

type StackError struct {
	Op  error
	err string
}

func NewStackError(Op error, err string) error {
	return StackError{Op: Op, err: err}
}

func (s StackError) Error() string {
	return s.err
}

func (s StackError) Unwrap() error {
	return s.Op
}

type TsStack interface {
	Tspush(x interface{}) error
	Tspop() (interface{}, error)
	Tssize() uint64
	Tsclear() error
	Tslook(size uint64) (interface{}, error)
}

type NoTsStack interface {
	Push(x interface{}) error
	Pop() (interface{}, error)
	Size() uint64
	Clear() error
	Look(size uint64) (interface{}, error)
}

type Stack interface {
	TsStack
	NoTsStack
}
