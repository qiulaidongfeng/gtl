package stack

import (
	"errors"
)

var (
	StackEmpty        error = errors.New("Stack,Empty")
	StackSizeExceeded error = errors.New("Stack size exceeded")
	StackPushFail     error = errors.New("Stack push fail")
	StackClearFail    error = errors.New("Stack clear fail")
)

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
