//本包实现了一个增强型的errors
//
//本包是实验性的，可能出现不兼容的修改
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

//错误原因类型
type TypeError string

//增强的error
type Error interface {
	//返回错误类型
	Type() TypeError
	//错误信息
	error
	//返回被包装的错误，nil表示没有被包装的错误
	Unwrap() Error
}

//Error接口的实现
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

//创建错误，返回的是GoError结构体
func NewError(Type TypeError, err string) Error {
	return &GoError{errorType: Type, err: err}
}
