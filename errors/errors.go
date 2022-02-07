// error
package errors

//错误接口
type Error interface {
	//go自带错误接口
	error
	//返回不带栈踪迹信息的方法
	ErrorNoStack() string
}

type Errorstring string

func (err Errorstring) Error() string {
	return string(err)
}

func (err Errorstring) ErrorNoStack() string {
	return string(err)
}

func New(err string) error {
	return Errorstring(err)
}
