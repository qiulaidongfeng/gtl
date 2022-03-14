package errors

import (
	"errors"
	"reflect"
)

//错误接口
type StackError interface {
	//go自带错误接口
	error
	//返回不带栈踪迹信息的方法
	ErrorNoStack() string
}

//带有包装错误的错误接口
type WrapErrorType interface {
	//go自带错误接口
	error
	//返回被包装错误的方法
	Unwrap() error
}

//增强版带有包装错误的错误接口
type PlusWrapError interface {
	//go自带错误接口
	error
	//返回被包装错误的方法
	Unwrap() error
	//返回不带栈踪迹信息的方法
	ErrorNoStack() string
}

type Errorstring string

func (err *Errorstring) Error() string {
	return string(*err)
}

func (err *Errorstring) ErrorNoStack() string {
	return string(*err)
}

func New(err string) error {
	errorerr := Errorstring(err)
	return &(errorerr)
}

func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

//比较没带有被包装错误的错误与compared（被比较的）错误是否相等
func TypeEqual(err, compared error) bool {
	errtype := reflect.TypeOf(err)
	comparedtype := reflect.TypeOf(compared)
	return errtype == comparedtype
}

/*
	比较err(带有被包装错误的错误)中的所有错误是否存在与compared（被比较的错误）相等的
	err需要实现WrapErrorType接口
	返回第一个相等的错误值的被包装的层数，0表示顶层
*/
func AllTypeEqual(err WrapErrorType, compared error) (int, bool) {
	n := 0
	ok := false
	for {
		errtype := reflect.TypeOf(err)
		comparedtype := reflect.TypeOf(compared)
		if errtype == comparedtype {
			return n, true
		}
		if err, ok = err.Unwrap().(WrapErrorType); ok == false {
			return n, false
		}
		n++
	}
}

//返回最后一层错误
func Cause(err error) error {
	ok := false
	var err1 WrapErrorType
	if err, ok := err.(interface{ Cause() error }); ok == true {
		return err.Cause()
	}
	for {
		if err1, ok = err.(WrapErrorType); ok == true {
			err = err1.Unwrap()
		} else {
			return err
		}
	}
}
