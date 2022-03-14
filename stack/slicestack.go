package stack

import (
	"sync"
	"sync/atomic"
)

type Slicestack struct {
	slice []interface{}
	size  *uint64
	mutex sync.RWMutex
}

func Newslicestack() Slicestack {
	s := Slicestack{
		slice: make([]interface{}, 0, 2),
		size:  new(uint64),
	}
	return s
}

func (s *Slicestack) Push(x interface{}) error {
	s.slice = append(s.slice[:((*s.size)+1)], x)
	*s.size++
	return nil
}

func (s *Slicestack) TsPush(x interface{}) error {
	s.mutex.Lock()
	s.slice = append(s.slice[:((*s.size)+1)], x)
	*s.size += 1
	s.mutex.Unlock()
	return nil
}

func (s *Slicestack) Pop() (x interface{}, err error) {
	if (*s.size) == 0 {
		err = StackEmpty
		return x, err
	}
	x = s.slice[*s.size]
	*s.size -= 1
	return x, nil
}

func (s *Slicestack) TsPop() (x interface{}, err error) {
	s.mutex.Lock()
	if *s.size == 0 {
		err = StackEmpty
		return x, err
	}
	x = s.slice[*s.size]
	*s.size -= 1
	s.mutex.Unlock()
	return x, nil
}

func (s *Slicestack) Size() uint64 {
	return *s.size
}

func (s *Slicestack) TsSize() uint64 {
	return atomic.LoadUint64(s.size)
}

func (s *Slicestack) Clear() error {
	s.slice = make([]interface{}, 0, 2)
	*s.size = 0
	return nil
}

func (s *Slicestack) TsClear() error {
	s.mutex.Lock()
	s.slice = make([]interface{}, 0, 2)
	*s.size = 0
	s.mutex.Unlock()
	return nil
}

func (s Slicestack) Look(size uint64) (interface{}, error) {
	if *s.size < size {
		return nil, StackSizeExceeded
	}
	return s.slice[size], nil
}

func (s Slicestack) TsLook(size uint64) (interface{}, error) {
	s.mutex.Unlock()
	if *s.size < size {
		return nil, StackSizeExceeded
	}
	s.mutex.RUnlock()
	return s.slice[size], nil
}
