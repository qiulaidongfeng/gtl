// wrap
package errors

//被包装的错误类型
type WrapError struct {
	err     error
	message string
}

func (err *WrapError) Error() string {
	return err.message + err.err.Error()
}

//返回被包装的错误
func (err *WrapError) Unwrap() error {
	return err.err
}

//包装错误
func NewWrapError(err error, message string) WrapErrorType {
	return &WrapError{err: err, message: message}
}
