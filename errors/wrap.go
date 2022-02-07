// warp
package errors

type WrapError struct {
	err     error
	message string
}

func (err WrapError) Error() string {
	return err.message + err.err.Error()
}

func (err WrapError) Unwrap() error {
	return err.err
}

func WithMessage(err error, message string) error {
	return WrapError{err: err, message: message}
}

func WrapError(err error, message string) error {
	return WrapError{err: err, message: message}
}
