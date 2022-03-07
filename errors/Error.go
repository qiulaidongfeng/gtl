// Error
package errors

const (
	//系统错误
	OsError TypeError = "Os , Error"
	//io错误
	IoError TypeError = "Io , Error"
	//网络错误
	NetworkError TypeError = "Network , Error"
	//并发错误
	SyncError TypeError = "Sync , Error"
	//时间错误
	TimeError TypeError = "Time , Error"
	//未知错误
	UnknownError TypeError = "Unknown , Error"
)

type TypeError string

type Error interface {
	//返回错误类型
	Type() TypeError
	//错误信息
	error
	//返回被包装的错误，nil表示没有被保装的错误
	Unwrap() Error
}

type GoError struct {
	errorType TypeError
	err       string
}

func (err *GoError) Error() string {
	return err.err
}

func (err *GoError) Unwrap() Error {
	return nil
}

func (err *GoError) Type() TypeError {
	return err.errorType
}

func NewError(Type TypeError, err string) Error {
	return &GoError{errorType: Type, err: err}
}
