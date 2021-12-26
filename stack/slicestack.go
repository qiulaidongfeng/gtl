package stack

import (
	"errors"
	"sync"
)

type slicestack struct {
	slice []interface{}
	size  uint64
	mutex sync.RWMutex
}

func Newslicestack() slicestack {
	s := slicestack{
		slice: make([]interface{}, 0, 2),
		size:  uint64(0),
	}
	return s
}

func (s *slicestack) Push(x interface{}) {
	s.slice = append(s.slice[:(s.size+1)], x)
	s.size++
	return
}

func (s *slicestack) Pop() (x interface{}, err error) {
	if s.size == 0 {
		err = errors.New("slicestack,Empty")
		return x, err
	}
	x = s.slice[s.size]
	s.size -= 1
	return x, nil

}

func (s *slicestack) Size() uint64 {
	return s.size
}

func (s *slicestack) Clear() {
	s.slice = make([]interface{}, 0, 2)
	s.size = 0
}

func (s slicestack) Look(size uint64) (interface{}, error) {
	if s.size < size {
		return nil, errors.New("Stack size exceeded")
	}
	return s.slice[size], nil
}
