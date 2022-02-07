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
	if ok == false {
		return nil
	}
	return u.Unwrap()
}

func Is(err, target error) bool {
	if target == nil {
		return err == target
	}
	for {
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		if err = Unwrap(err); err == nil {
			return false
		}
	}
}
