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
	pushpp int64 = -1
	poppp  int64 = 1
)

const (
	Poptime  = int64(70)
	Pushtime = int64(140)
	Waittime = int64(100)
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
	slice    []int8
	size     uint64
	scap     uint64
	pushn    int64
	popn     int64
	pp       int64
	pushtime int64
	poptime  int64
	waittime int64
	mutex    sync.RWMutex
}

type GLMstack = GenericLowMemoryStack

func NewGLMstack() GLMstack {
	s := GLMstack{
		slice:    make([]int8, 2, 2),
		size:     2,
		scap:     2,
		pushn:    0,
		popn:     0,
		pp:       0,
		pushtime: Pushtime,
		poptime:  Poptime,
		waittime: Waittime,
	}
	go s.pprecode()
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
	for i := uint64(0); i < s.scap; i++ {
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
	for i := uint64(0); i < s.scap; i++ {
		nslice[i] = s.slice[i]
	}
	s.slice = nslice
	s.mutex.Unlock()
	return
}

/*//记录现在是入栈还是出栈
func (s *GLMstack) pprecode() {
	for {
		bol := atomic.CompareAndSwapInt64(&s.pp, poppp, poppp) //现在是出栈
		if bol == true {
			popn := atomic.LoadInt64(&s.popn)
			if popn == 0 { //现在没有实际出栈
				atomic.SwapInt64(&s.pp, 0)
				popnold := atomic.LoadInt64(&s.popn)
				if popnold != 0 { //出现实际出栈
					atomic.SwapInt64(&s.pp, poppp) //记录为出栈
					time.Sleep(time.Duration(s.pushtime * s.pushn))
					continue
				} else {
					continue
				}
			}
		} else {
			bol := atomic.CompareAndSwapInt64(&s.pp, pushpp, pushpp) //现在是入栈
			if bol == true {
				pushn := atomic.LoadInt64(&s.pushn)
				if pushn == 0 { //现在没有实际入栈
					atomic.SwapInt64(&s.pp, 0)
					pushnold := atomic.LoadInt64(&s.pushn)
					if pushnold != 0 { //出现实际入栈
						atomic.StoreInt64(&s.pp, pushpp) //记录为入栈
						time.Sleep(time.Duration(s.poptime * s.popn))
						continue
					}
				} else {
					continue
				}
			} else { //现在没有栈操作
				time.Sleep(time.Duration(s.waittime))
				continue
			}
		}
	}
}
*/

//记录现在是入栈还是出栈
func (s *GLMstack) pprecode() {
	for {
		pp := atomic.LoadInt64(&s.pp)
		if pp == 0 {
			time.Sleep(time.Duration(s.waittime))
		} else if pp == poppp {
			popn := atomic.LoadInt64(&s.popn)
			if popn == 0 {
				atomic.SwapInt64(&s.pp, 0)
			} else {
				time.Sleep(time.Duration(s.poptime * popn))
			}
		} else {
			pushn := atomic.LoadInt64(&s.pushn)
			if pushn == 0 {
				atomic.SwapInt64(&s.pp, 0)
			} else {
				time.Sleep(time.Duration(s.pushtime * pushn))
			}
		}
	}
}

//记录为入栈时操作
func (s *GLMstack) pushrecord() {
	atomic.AddInt64(&s.pushn, 1)
	for {
		rw := atomic.LoadInt64(&s.pp)
		if rw == -1 { //正在入栈操作
			return
		} else if rw == 0 {
			bol := atomic.CompareAndSwapInt64(&s.pp, 0, pushpp)
			if bol == true { //无操作
				return
			} else { //正在出栈操作
				time.Sleep(time.Duration(s.poptime * s.popn))
			}
		} else { //正在出栈操作
			time.Sleep(time.Duration(s.poptime * s.popn))
		}
	}
}

func (s *GLMstack) pushok() {
	atomic.AddInt64(&s.pushn, -1)
}

func (s *GLMstack) popok() {
	atomic.AddInt64(&s.popn, -1)
}

//记录为出栈时操作
func (s *GLMstack) poprecord() {
	atomic.AddInt64(&s.popn, 1)
	for {
		rw := atomic.LoadInt64(&s.pp)
		if rw == 1 { //正在出栈操作
			return
		} else if rw == 0 {
			bol := atomic.CompareAndSwapInt64(&s.pp, 0, poppp)
			if bol == true { //无操作
				return
			} else {
				time.Sleep(time.Duration(s.pushtime * s.pushn))
			}
		} else {
			time.Sleep(time.Duration(s.pushtime * s.pushn))
		}
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
	s.pushrecord() //入栈记录
	if s.size+size >= s.scap {
		s.mutex.RUnlock()
		s.scap = s.addcap(s.size + size)
		s.mutex.RLock()
	}
	sizeold := atomic.AddUint64(&s.size, size)
	sizeold = atomic.AddUint64(&sizeold, ^uint64(s.size-1))
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
	for i := uint64(0); i < size; i++ {
		*(*int8)(unsafe.Pointer(vptr + (uintptr(i)))) = *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
	}
	return nil
}

func (s *GLMstack) TsPopptr(ptr *unsafe.Pointer, size uint64) error {
	s.mutex.RLock()
	s.poprecord() //出栈记录
	sizei := atomic.AddUint64(&s.size, ^(size - 1))
	vptr := uintptr(unsafe.Pointer(&s.slice[0])) + uintptr(sizei)
	uptr := uintptr(*ptr)
	for i := uint64(0); i < size; i++ { //实际出栈
		*(*int8)(unsafe.Pointer(vptr + (uintptr(i)))) = *(*int8)(unsafe.Pointer(uptr + uintptr(i)))
	}
	s.popok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) TsLoadpp() int64 {
	return atomic.LoadInt64(&s.pp)
}

func (s *GLMstack) TsLoadpopn() int64 {
	return atomic.LoadInt64(&s.popn)
}

func (s *GLMstack) TsLoadpushn() int64 {
	return atomic.LoadInt64(&s.pushn)
}
