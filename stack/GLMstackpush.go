// GLMstackpush
package stack

import (
	"reflect"
	"unsafe"
)

func (s *GLMstack) Push(x interface{}) error {
	kind := reflect.ValueOf(x).Kind()
	if kind >= reflect.Int && kind <= reflect.Int64 {
		switch Type := x.(type) {
		case int:
			err := s.Pushint(Type)
			if err != nil {
				return err
			}
		case int64:
			err := s.Pushint64(Type)
			if err != nil {
				return err
			}
		case int32:
			err := s.Pushint32(Type)
			if err != nil {
				return err
			}
		case int8:
			err := s.Pushint8(Type)
			if err != nil {
				return err
			}
		case int16:
			err := s.Pushint16(Type)
			if err != nil {
				return err
			}
		}
	} else if kind == reflect.Bool {
		Type := x.(bool)
		s.PushBool(Type)
	} else if kind >= reflect.Uint && kind <= reflect.Uint64 {
		switch Type := x.(type) {
		case uint:
			err := s.Pushuint(Type)
			if err != nil {
				return err
			}
		case uint64:
			err := s.Pushuint64(Type)
			if err != nil {
				return err
			}
		case uint32:
			err := s.Pushuint32(Type)
			if err != nil {
				return err
			}
		case uint8:
			err := s.Pushuint8(Type)
			if err != nil {
				return err
			}
		case uint16:
			err := s.Pushuint16(Type)
			if err != nil {
				return err
			}
		}
	} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
		switch Type := x.(type) {
		case float32:
			err := s.PushFloat32(Type)
			if err != nil {
				return err
			}
		case float64:
			err := s.PushFloat64(Type)
			if err != nil {
				return err
			}
		}
	} else if kind == reflect.Interface {
		s.PushInterface(x)
	} else if kind == reflect.Uintptr {
		Type := x.(uintptr)
		err := s.PushUintptr(Type)
		if err != nil {
			return err
		}
	} else if kind >= reflect.Complex64 && kind < reflect.Complex128 {
		switch Type := x.(type) {
		case complex64:
			err := s.PushComplex64(Type)
			if err != nil {
				return err
			}
		case complex128:
			err := s.PushComplex128(Type)
			if err != nil {
				return err
			}
		}
	} else {
		return StackPushFail
	}
	return StackPushFail
}

func (s *GLMstack) TsPush(x interface{}) error {
	kind := reflect.ValueOf(x).Kind()
	if kind >= reflect.Int && kind <= reflect.Int64 {
		switch Type := x.(type) {
		case int:
			err := s.TsPushint(Type)
			if err != nil {
				return err
			}
		case int64:
			err := s.TsPushint64(Type)
			if err != nil {
				return err
			}
		case int32:
			err := s.TsPushint32(Type)
			if err != nil {
				return err
			}
		case int8:
			err := s.TsPushint8(Type)
			if err != nil {
				return err
			}
		case int16:
			err := s.TsPushint16(Type)
			if err != nil {
				return err
			}
		}
	} else if kind == reflect.Bool {
		Type := x.(bool)
		s.TsPushBool(Type)
	} else if kind >= reflect.Uint && kind <= reflect.Uint64 {
		switch Type := x.(type) {
		case uint:
			err := s.TsPushuint(Type)
			if err != nil {
				return err
			}
		case uint64:
			err := s.TsPushuint64(Type)
			if err != nil {
				return err
			}
		case uint32:
			err := s.TsPushuint32(Type)
			if err != nil {
				return err
			}
		case uint8:
			err := s.TsPushuint8(Type)
			if err != nil {
				return err
			}
		case uint16:
			err := s.TsPushuint16(Type)
			if err != nil {
				return err
			}
		}
	} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
		switch Type := x.(type) {
		case float32:
			err := s.TsPushFloat32(Type)
			if err != nil {
				return err
			}
		case float64:
			err := s.TsPushFloat64(Type)
			if err != nil {
				return err
			}
		}
	} else if kind == reflect.Interface {
		s.TsPushInterface(x)
	} else if kind == reflect.Uintptr {
		Type := x.(uintptr)
		err := s.TsPushUintptr(Type)
		if err != nil {
			return err
		}
	} else if kind >= reflect.Complex64 && kind < reflect.Complex128 {
		switch Type := x.(type) {
		case complex64:
			err := s.TsPushComplex64(Type)
			if err != nil {
				return err
			}
		case complex128:
			err := s.TsPushComplex128(Type)
			if err != nil {
				return err
			}
		}
	} else {
		return StackPushFail
	}
	return StackPushFail
}

func (s *GLMstack) Pushptr(ptr unsafe.Pointer, size uint64) error {
	safe := s.pushsafetycheck(size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + size)
	}
	uptr := uintptr(ptr)
	for i := uint64(0); i < size; i++ {
		size := s.size + i
		value := *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
		s.slice[size] = value
	}
	s.size += size
	return nil
}

func (s *GLMstack) TsPushptr(ptr unsafe.Pointer, size uint64) error {
	s.mutex.RLock()
	s.pushrecord()                       //入栈记录
	sizeold := s.tspushsafetycheck(size) //入栈安全检查
	uptr := uintptr(ptr)
	for i := uint64(0); i < size; i++ { //实际入栈
		sizei := sizeold + i
		value := *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
		s.slice[sizei] = value
	}
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}
