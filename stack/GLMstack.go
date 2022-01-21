// gsstack
package stack

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	Int8size       uint64 = uint64((unsafe.Sizeof(int8(1))))
	Int16size      uint64 = uint64((unsafe.Sizeof(int16(1))))
	Int32size      uint64 = uint64((unsafe.Sizeof(int32(1))))
	Int64size      uint64 = uint64((unsafe.Sizeof(int64(1))))
	Intsize        uint64 = uint64((unsafe.Sizeof(int(1))))
	Uint8size      uint64 = uint64((unsafe.Sizeof(uint8(1))))
	Uint16size     uint64 = uint64((unsafe.Sizeof(uint16(1))))
	Uint32size     uint64 = uint64((unsafe.Sizeof(uint32(1))))
	Uint64size     uint64 = uint64((unsafe.Sizeof(uint64(1))))
	Uintsize       uint64 = uint64((unsafe.Sizeof(uint(1))))
	Bytesize       uint64 = uint64((unsafe.Sizeof(byte(1))))
	Runesize       uint64 = uint64((unsafe.Sizeof(rune(1))))
	Boolsize       uint64 = uint64((unsafe.Sizeof(bool(true))))
	Float32size    uint64 = uint64((unsafe.Sizeof(float32(1.0))))
	Float64size    uint64 = uint64((unsafe.Sizeof(float64(2.0))))
	Uintptrsize    uint64 = uint64((unsafe.Sizeof(uintptr(7))))
	Complex64size  uint64 = uint64((unsafe.Sizeof(complex64(6 + 9i))))
	Complex128size uint64 = uint64((unsafe.Sizeof(complex128(8 + 9i))))
	Interfacesize  uint64 = uint64((unsafe.Sizeof(*(new(interface{})))))
)

type GenericLowMemoryStack struct {
	slice []int8
	size  uint64
	scap  uint64
	rw    int64
	n     int64
	mutex sync.RWMutex
}

type GLMstack = GenericLowMemoryStack

func NewGLMstack() GLMstack {
	s := GLMstack{
		slice: make([]int8, 2, 2),
		size:  2,
		scap:  2,
		rw:    0,
	}
	return s
}

func (s *GLMstack) addcap(size uint64) (ncap uint64) {
	ncap = uint64(cap(s.slice))
	for ncap <= size {
		if ncap < 1024 {
			ncap += ncap
		} else {
			ncap += ncap / 4
		}
	}
	nslice := make([]int8, ncap, ncap)
	// nptr := uintptr(unsafe.Pointer(&nslice[0]))
	// uptr := uintptr(unsafe.Pointer(&s.slice[0]))
	for i := uint64(0); i < s.size; i++ {
		nslice[i] = s.slice[i]
	}
	s.slice = nslice
	return
}

func (s *GLMstack) Tsaddcap(size uint64) (ncap uint64) {
	s.mutex.Lock()
	ncap = uint64(cap(s.slice))
	for ncap <= size {
		if ncap < 1024 {
			ncap += ncap
		} else {
			ncap += ncap / 4
		}
	}
	nslice := make([]int8, ncap, ncap)
	for i := uint64(0); i < s.size; i++ {
		nslice[i] = s.slice[i]
	}
	s.slice = nslice
	s.mutex.Unlock()
	return
}

func (s *GLMstack) rwrecord(rw int, waittime time.Duration) {
	if rw == 1 {
		for {
			bol := atomic.CompareAndSwapInt64(&s.rw, -1, -1)
			if bol == true {
				time.Sleep(waittime)
				continue
			} else {
				old := atomic.SwapInt64(&s.rw, 1)
				if old == -1 {
					atomic.StoreInt64(&s.rw, -1)
					continue
				}
				atomic.AddInt64(&s.n, 1)
			}
		}
	} else if rw == -1 {
		for {
			bol := atomic.CompareAndSwapInt64(&s.rw, 1, 1)
			if bol == true {
				time.Sleep(waittime)
				continue
			} else {
				old := atomic.SwapInt64(&s.rw, -1)
				if old == -1 {
					atomic.StoreInt64(&s.rw, 1)
					continue
				}
				atomic.AddInt64(&s.n, 1)
			}
		}
	} else {
		panic("err:无效操作")
	}
}

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
	if s.size+size >= s.scap {
		s.scap = s.addcap(s.size + size)
	}
	for i := uint64(0); i < size; i++ {
		size := s.size + i
		value := *(*int8)(unsafe.Pointer(uintptr(ptr) + uintptr(i)))
		s.slice[size] = value
	}
	s.size += size
	return nil
}

func (s *GLMstack) TsPushptr(ptr unsafe.Pointer, size uint64) error {
	s.mutex.RLock()
	if s.size+size >= s.scap {
		s.mutex.RUnlock()
		s.scap = s.addcap(s.size + size)
		s.mutex.RLock()
	}
	sizeold := atomic.AddUint64(&s.size, size)
	sizeold = atomic.AddUint64(&sizeold, ^uint64(s.size-1))
	for i := uint64(0); i < size; i++ {
		sizei := sizeold + i
		value := *(*int8)(unsafe.Pointer(uintptr(ptr) + uintptr(i)))
		s.slice[sizei] = value
	}
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Size() uint64 {
	return s.size
}

func (s *GLMstack) Tssize() uint64 {
	return atomic.LoadUint64(&s.size)
}

func (s *GLMstack) Clear() error {
	s.scap = 2
	s.size = 2
	s.slice = make([]int8, 2, 2)
	return nil
}

func (s *GLMstack) Tsclear() error {
	s.mutex.Lock()
	s.scap = 2
	s.size = 2
	s.slice = make([]int8, 2, 2)
	s.mutex.Unlock()
	return nil
}

func (s *GLMstack) Popptr(ptr *unsafe.Pointer, size uint64) error {
	sizei := s.size
	s.size -= size
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	uptr := uintptr(*ptr)
	for i := uint64(0); i < s.size; i++ {
		*(*int8)(unsafe.Pointer(uptr + uintptr(i))) = *(*int8)(unsafe.Pointer(vptr + (uintptr(i))))
	}
	return nil
}

func (s *GLMstack) TsPopptr(ptr *unsafe.Pointer, size uint64) error {
	s.mutex.RLock()
	v := make([]int8, size, size)
	sizei := atomic.AddUint64(&s.size, ^(size - 1))
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	for i := uint64(0); i < size; i++ {
		v[i] = *(*int8)(unsafe.Pointer(vptr + (uintptr(i))))
	}
	*ptr = unsafe.Pointer(&v[0])
	s.mutex.RUnlock()
	return nil
}
