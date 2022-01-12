package stack

import (
	"sync/atomic"
	"unsafe"
)

func (s *GLMstack) Pushint(x int) error {
	if (*s.size)+Intsize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Intsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int)(sp)
	*sp2 = x
	*s.size += Intsize
	return nil
}

func (s *GLMstack) TsPushint(x int) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Intsize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Intsize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint8(x int8) error {
	if (*s.size)+Int8size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Int8size)
	}
	s.slice[*s.size] = x
	*s.size++
	return nil
}

func (s *GLMstack) TsPushint8(x int8) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int8size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Int8size)
		s.mutex.RLock()
	}
	s.slice[nsize] = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint16(x int16) error {
	if (*s.size)+Int16size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Int16size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	*s.size += Int16size
	return nil
}

func (s *GLMstack) TsPushint16(x int16) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int16size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Int16size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint32(x int32) error {
	if (*s.size)+Int32size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Int32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int32)(sp)
	*sp2 = x
	*s.size += Int32size
	return nil
}

func (s *GLMstack) TsPushint32(x int32) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int32size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Int32size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int32)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushint64(x int64) error {
	if (*s.size)+Int64size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Int64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int64)(sp)
	*sp2 = x
	*s.size += Int64size
	return nil
}

func (s *GLMstack) TsPushint64(x int64) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int64size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Int64size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int64)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint(x uint) error {
	if (*s.size)+Uintsize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Uintsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint)(sp)
	*sp2 = x
	*s.size += Uintsize
	return nil
}

func (s *GLMstack) TsPushuint(x uint) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uintsize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Uintsize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint8(x uint8) error {
	if (*s.size)+Uint8size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Uint8size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint8)(sp)
	*sp2 = x
	*s.size += Uint8size
	return nil
}

func (s *GLMstack) TsPushuint8(x uint8) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint8size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Uint8size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint8)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint16(x uint16) error {
	if (*s.size)+Uint16size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Uint16size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint16)(sp)
	*sp2 = x
	*s.size += Uint16size
	return nil
}

func (s *GLMstack) TsPushuint16(x uint16) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint16size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Uint16size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint16)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint32(x uint32) error {
	if (*s.size)+Uint32size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Uint32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint32)(sp)
	*sp2 = x
	*s.size += Uint32size
	return nil
}

func (s *GLMstack) TsPushuint32(x uint32) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint32size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Uint32size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint32)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) Pushuint64(x uint64) error {
	if (*s.size)+Uint64size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Uint64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint64)(sp)
	*sp2 = x
	*s.size += Uint64size
	return nil
}

func (s *GLMstack) TsPushuint64(x uint64) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint64size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Uint64size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint64)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushByte(x byte) error {
	if (*s.size)+Bytesize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Bytesize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*byte)(sp)
	*sp2 = x
	*s.size += Bytesize
	return nil
}

func (s *GLMstack) TsPushByte(x byte) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Bytesize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Bytesize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*byte)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushRune(x rune) error {
	if (*s.size)+Runesize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Runesize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*rune)(sp)
	*sp2 = x
	*s.size += Runesize
	return nil
}

func (s *GLMstack) TsPushRune(x rune) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Runesize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Runesize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*rune)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushBool(x bool) error {
	if (*s.size)+Boolsize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Boolsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*bool)(sp)
	*sp2 = x
	*s.size += Runesize
	return nil
}

func (s *GLMstack) TsPushBool(x bool) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Boolsize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Boolsize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*bool)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushFloat32(x float32) error {
	if (*s.size)+Float32size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Float32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float32)(sp)
	*sp2 = x
	*s.size += Float32size
	return nil
}

func (s *GLMstack) TsPushFloat32(x float32) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Float32size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Float32size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float32)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushFloat64(x float64) error {
	if (*s.size)+Float64size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Float64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float64)(sp)
	*sp2 = x
	*s.size += Float64size
	return nil
}

func (s *GLMstack) TsPushFloat64(x float64) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Float64size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Float64size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*float64)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushUintptr(x uintptr) error {
	if (*s.size)+Uintptrsize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Uintptrsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uintptr)(sp)
	*sp2 = x
	*s.size += Uintptrsize
	return nil
}

func (s *GLMstack) TsPushUintptr(x uintptr) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uintptrsize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Uintptrsize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uintptr)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushComplex128(x complex128) error {
	if (*s.size)+Complex128size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Complex128size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex128)(sp)
	*sp2 = x
	*s.size += Complex128size
	return nil
}

func (s *GLMstack) TsPushComplex128(x complex128) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Complex128size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Complex128size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex128)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushComplex64(x complex64) error {
	if (*s.size)+Complex64size >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Complex64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex64)(sp)
	*sp2 = x
	*s.size += Complex64size
	return nil
}

func (s *GLMstack) TsPushComplex64(x complex64) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Complex64size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Complex64size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*complex64)(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}

func (s *GLMstack) PushInterface(x interface{}) error {
	if (*s.size)+Interfacesize >= (*s.scap) {
		*s.scap = s.addcap((*s.size) + Interfacesize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*interface{})(sp)
	*sp2 = x
	*s.size += Interfacesize
	return nil
}

func (s *GLMstack) TsPushInterface(x interface{}) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Interfacesize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap((*s.size) + Interfacesize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*interface{})(sp)
	*sp2 = x
	s.mutex.RUnlock()
	return nil
}
