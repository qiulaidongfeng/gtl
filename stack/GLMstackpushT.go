package stack

import (
	"unsafe"
)

func (s *GLMstack) Pushint(x int) error {
	safe := s.pushsafetycheck(Intsize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Intsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int)(sp)
	*sp2 = x
	s.size += Intsize
	return nil
}

func (s *GLMstack) TsPushint(x int) error {
	s.mutex.RLock()
	s.pushrecord()                     //入栈记录
	sl := s.tspushsafetycheck(Intsize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint8(x int8) error {
	safe := s.pushsafetycheck(Int8size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Int8size)
	}
	s.slice[s.size] = x
	s.size++
	return nil
}

func (s *GLMstack) TsPushint8(x int8) error {
	s.mutex.RLock()
	s.pushrecord()                         //入栈记录
	nsize := s.tspushsafetycheck(Int8size) //入栈安全检查
	s.slice[nsize] = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint16(x int16) error {
	safe := s.pushsafetycheck(Int16size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Int16size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	s.size += Int16size
	return nil
}

func (s *GLMstack) TsPushint16(x int16) error {
	s.mutex.RLock()
	s.pushrecord()                       //入栈记录
	sl := s.tspushsafetycheck(Int16size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint32(x int32) error {
	safe := s.pushsafetycheck(Int32size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Int32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int32)(sp)
	*sp2 = x
	s.size += Int32size
	return nil
}

func (s *GLMstack) TsPushint32(x int32) error {
	s.mutex.RLock()
	s.pushrecord()                       //入栈记录
	sl := s.tspushsafetycheck(Int32size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int32)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint64(x int64) error {
	safe := s.pushsafetycheck(Int64size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Int64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int64)(sp)
	*sp2 = x
	s.size += Int64size
	return nil
}

func (s *GLMstack) TsPushint64(x int64) error {
	s.mutex.RLock()
	s.pushrecord()                       //入栈记录
	sl := s.tspushsafetycheck(Int64size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int64)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint(x uint) error {
	safe := s.pushsafetycheck(Uintsize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Uintsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint)(sp)
	*sp2 = x
	s.size += Uintsize
	return nil
}

func (s *GLMstack) TsPushuint(x uint) error {
	s.mutex.RLock()
	s.pushrecord()                      //入栈记录
	sl := s.tspushsafetycheck(Uintsize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint8(x uint8) error {
	safe := s.pushsafetycheck(Uint8size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Uint8size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint8)(sp)
	*sp2 = x
	s.size += Uint8size
	return nil
}

func (s *GLMstack) TsPushuint8(x uint8) error {
	s.mutex.RLock()
	s.pushrecord()                       //入栈记录
	sl := s.tspushsafetycheck(Uint8size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint8)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint16(x uint16) error {
	safe := s.pushsafetycheck(Uint16size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Uint16size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint16)(sp)
	*sp2 = x
	s.size += Uint16size
	return nil
}

func (s *GLMstack) TsPushuint16(x uint16) error {
	s.mutex.RLock()
	s.pushrecord()                        //入栈记录
	sl := s.tspushsafetycheck(Uint16size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint16)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint32(x uint32) error {
	safe := s.pushsafetycheck(Uint32size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Uint32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint32)(sp)
	*sp2 = x
	s.size += Uint32size
	return nil
}

func (s *GLMstack) TsPushuint32(x uint32) error {
	s.mutex.RLock()
	s.pushrecord()                        //入栈记录
	sl := s.tspushsafetycheck(Uint32size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint32)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint64(x uint64) error {
	safe := s.pushsafetycheck(Uint64size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Uint64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint64)(sp)
	*sp2 = x
	s.size += Uint64size
	return nil
}

func (s *GLMstack) TsPushuint64(x uint64) error {
	s.mutex.RLock()
	s.pushrecord()                        //入栈记录
	sl := s.tspushsafetycheck(Uint64size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint64)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushByte(x byte) error {
	safe := s.pushsafetycheck(Bytesize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Bytesize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*byte)(sp)
	*sp2 = x
	s.size += Bytesize
	return nil
}

func (s *GLMstack) TsPushByte(x byte) error {
	s.mutex.RLock()
	s.pushrecord()                      //入栈记录
	sl := s.tspushsafetycheck(Bytesize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*byte)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushRune(x rune) error {
	safe := s.pushsafetycheck(Runesize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Runesize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*rune)(sp)
	*sp2 = x
	s.size += Runesize
	return nil
}

func (s *GLMstack) TsPushRune(x rune) error {
	s.mutex.RLock()
	s.pushrecord()                      //入栈记录
	sl := s.tspushsafetycheck(Runesize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*rune)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushBool(x bool) error {
	safe := s.pushsafetycheck(Boolsize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Boolsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*bool)(sp)
	*sp2 = x
	s.size += Runesize
	return nil
}

func (s *GLMstack) TsPushBool(x bool) error {
	s.mutex.RLock()
	s.pushrecord()                      //入栈记录
	sl := s.tspushsafetycheck(Boolsize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*bool)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushFloat32(x float32) error {
	safe := s.pushsafetycheck(Float32size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Float32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float32)(sp)
	*sp2 = x
	s.size += Float32size
	return nil
}

func (s *GLMstack) TsPushFloat32(x float32) error {
	s.mutex.RLock()
	s.pushrecord()                         //入栈记录
	sl := s.tspushsafetycheck(Float32size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float32)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushFloat64(x float64) error {
	safe := s.pushsafetycheck(Float64size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Float64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float64)(sp)
	*sp2 = x
	s.size += Float64size
	return nil
}

func (s *GLMstack) TsPushFloat64(x float64) error {
	s.mutex.RLock()
	s.pushrecord()                         //入栈记录
	sl := s.tspushsafetycheck(Float64size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float64)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushUintptr(x uintptr) error {
	safe := s.pushsafetycheck(Uintptrsize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Uintptrsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uintptr)(sp)
	*sp2 = x
	s.size += Uintptrsize
	return nil
}

func (s *GLMstack) TsPushUintptr(x uintptr) error {
	s.mutex.RLock()
	s.pushrecord()                         //入栈记录
	sl := s.tspushsafetycheck(Uintptrsize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uintptr)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushComplex128(x complex128) error {
	safe := s.pushsafetycheck(Complex128size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Complex128size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex128)(sp)
	*sp2 = x
	s.size += Complex128size
	return nil
}

func (s *GLMstack) TsPushComplex128(x complex128) error {
	s.mutex.RLock()
	s.pushrecord()                            //入栈记录
	sl := s.tspushsafetycheck(Complex128size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex128)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushComplex64(x complex64) error {
	safe := s.pushsafetycheck(Complex64size) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Complex64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex64)(sp)
	*sp2 = x
	s.size += Complex64size
	return nil
}

func (s *GLMstack) TsPushComplex64(x complex64) error {
	s.mutex.RLock()
	s.pushrecord()                           //入栈记录
	sl := s.tspushsafetycheck(Complex64size) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex64)(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushInterface(x interface{}) error {
	safe := s.pushsafetycheck(Interfacesize) //入栈安全检查
	if safe != nil {
		s.scap = s.addcap(s.size + Interfacesize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := s.size
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*interface{})(sp)
	*sp2 = x
	s.size += Interfacesize
	return nil
}

func (s *GLMstack) TsPushInterface(x interface{}) error {
	s.mutex.RLock()
	s.pushrecord()                           //入栈记录
	sl := s.tspushsafetycheck(Interfacesize) //入栈安全检查
	sp := unsafe.Pointer(&(s.slice[0]))
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*interface{})(sp)
	*sp2 = x
	s.pushok() //结束记录
	s.mutex.RUnlock()
	return nil
}
