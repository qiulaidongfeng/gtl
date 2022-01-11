package stack

import (
	"sync/atomic"
	"unsafe"
)

func (s *Sstack) Pushint(x int) error {
	if (*s.size)+Intsize >= (*s.scap) {
		*s.scap = s.addcap(Intsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int)(sp)
	*sp2 = x
	*s.size += Intsize
	return nil
}

func (s *Sstack) TsPushint(x int) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Intsize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Intsize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int)(sp)
	*sp2 = x
	*s.size += Intsize
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushint8(x int8) error {
	if (*s.size)+Int8size >= (*s.scap) {
		*s.scap = s.addcap(Int8size)
	}
	s.slice[*s.size] = x
	*s.size++
	return nil
}

func (s *Sstack) TsPushint8(x int8) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int8size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Int8size)
		s.mutex.RLock()
	}
	s.slice[nsize] = x
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushint16(x int16) error {
	if (*s.size)+Int16size >= (*s.scap) {
		*s.scap = s.addcap(Int16size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	*s.size += Int16size
	return nil
}

func (s *Sstack) TsPushint16(x int16) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int16size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Int16size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int16)(sp)
	*sp2 = x
	*s.size += Int16size
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushint32(x int32) error {
	if (*s.size)+Int32size >= (*s.scap) {
		*s.scap = s.addcap(Int32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int32)(sp)
	*sp2 = x
	*s.size += Int32size
	return nil
}

func (s *Sstack) TsPushint32(x int32) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int32size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Int32size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int32)(sp)
	*sp2 = x
	*s.size += Int32size
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushint64(x int64) error {
	if (*s.size)+Int64size >= (*s.scap) {
		*s.scap = s.addcap(Int64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int64)(sp)
	*sp2 = x
	*s.size += Int64size
	return nil
}

func (s *Sstack) TsPushint64(x int64) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Int64size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Int64size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*int64)(sp)
	*sp2 = x
	*s.size += Int64size
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushuint(x uint) error {
	if (*s.size)+Uintsize >= (*s.scap) {
		*s.scap = s.addcap(Uintsize)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint)(sp)
	*sp2 = x
	*s.size += Uintsize
	return nil
}

func (s *Sstack) TsPushuint(x uint) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uintsize)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Uintsize)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint)(sp)
	*sp2 = x
	*s.size += Uintsize
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushuint8(x uint8) error {
	if (*s.size)+Uint8size >= (*s.scap) {
		*s.scap = s.addcap(Uint8size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint8)(sp)
	*sp2 = x
	*s.size += Uint8size
	return nil
}

func (s *Sstack) TsPushuint8(x uint8) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint8size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Uint8size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint8)(sp)
	*sp2 = x
	*s.size += Uint8size
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushuint16(x uint16) error {
	if (*s.size)+Uint16size >= (*s.scap) {
		*s.scap = s.addcap(Uint16size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint16)(sp)
	*sp2 = x
	*s.size += Uint16size
	return nil
}

func (s *Sstack) TsPushuint16(x uint16) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint16size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Uint16size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint16)(sp)
	*sp2 = x
	*s.size += Uint16size
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushuint32(x uint32) error {
	if (*s.size)+Uint32size >= (*s.scap) {
		*s.scap = s.addcap(Uint32size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint32)(sp)
	*sp2 = x
	*s.size += Uint32size
	return nil
}

func (s *Sstack) TsPushuint32(x uint32) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint32size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Uint32size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint32)(sp)
	*sp2 = x
	*s.size += Uint32size
	s.mutex.RUnlock()
	return nil
}

func (s *Sstack) Pushuint64(x uint64) error {
	if (*s.size)+Uint64size >= (*s.scap) {
		*s.scap = s.addcap(Uint64size)
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint64)(sp)
	*sp2 = x
	*s.size += Uint64size
	return nil
}

func (s *Sstack) TsPushuint64(x uint64) error {
	s.mutex.RLock()
	nsize := atomic.AddUint64(s.size, Uint64size)
	if nsize >= *s.scap {
		s.mutex.RUnlock()
		*s.scap = s.Tsaddcap(Uint64size)
		s.mutex.RLock()
	}
	sp := unsafe.Pointer(&(s.slice[0]))
	sl := (*s.size)
	sp = unsafe.Pointer(uintptr(sp) + uintptr(sl))
	sp2 := (*uint64)(sp)
	*sp2 = x
	*s.size += Uint64size
	s.mutex.RUnlock()
	return nil
}
