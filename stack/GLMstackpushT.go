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
