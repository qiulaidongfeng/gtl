// slist
package list

import (
	"errors"
	"sync"
)

type Element struct {
	value   interface{}
	NextOne *Element
	mutex   sync.RWMutex
}

func (e *Element) Set(x interface{}) error {
	e.value = x
	return nil
}

func (e *Element) Get() (interface{}, error) {
	return e.value, nil
}

func (e *Element) Next() *Node {
	return e.NextOne
}

func (e *Element) Prev() *Node {
	return nil
}

func (e *Element) Tsset(x interface{}) error {
	e.mutex.Lock()
	e.value = x
	e.mutex.Unlock()
	return nil
}

func (e *Element) Tsget() (x interface{}, err error) {
	e.mutex.RLock()
	x = e.value
	e.mutex.RUnlock()
	return
}

func (e *Element) Tsnext() (x *Node) {
	e.mutex.RLock()
	x = e.NextOne
	e.mutex.RUnlock()
	return
}

func (e *Element) Tsprev() (x *Node) {
	x = nil
	return
}

type Slist struct {
	node Element
	len  uint64
}

func NewSlist() Slist {
	n := Slist{
		node: nil,
		len:  0,
	}
	return n
}

func (s *Slist) Lnsert(x Node) error {
	if s.len == 0 {
		s.node = &x
		return nil
	}
	var (
		next *Element = nil
	)
	for i := 0; i < (s.len - 1); i++ {
		next = &(s.node.NextOne)
	}
	*next = &x
	return nil

}

func (s *Slist) Remove() error {
	if s.len == 0 {
		return errors.New("Slist,Empty")
	}
	var (
		next *Element = nil
	)
	for i := 0; i < (s.len - 2); i++ {
		next = &(s.node.NextOne)
	}
	*next = nil
	return nil
}

func (s *Slist) Get(size uint64) (Node, error) {

}
