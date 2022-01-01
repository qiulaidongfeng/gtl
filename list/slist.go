// slist
package list

import (
	"errors"
	"fmt"
	"strings"
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

func (e *Element) Next() *Element {
	return e.NextOne
}

func (e *Element) Prev() *Element {
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

func (e *Element) Tsnext() (x *Element) {
	e.mutex.RLock()
	x = e.NextOne
	e.mutex.RUnlock()
	return
}

func (e *Element) Tsprev() (x *Element) {
	x = nil
	return
}

func (e *Element) String() string {
	return fmt.Sprintf("%s", e.value)
}

type SingleLinkedList struct {
	node *Element
	len  uint64
}

type Slist = SingleLinkedList

func NewSlist() Slist {
	n := Slist{
		len: 0,
	}
	return n
}

func (s *Slist) Lnsert(x *Element) error {
	if s.len == 0 {
		s.len++
		s.node = x
		return nil
	}
	if s.len == 1 {
		s.len++
		s.node.NextOne = x
		return nil
	}
	var next *Element
	for i := uint64(0); i <= (s.len - 2); i++ {
		next = s.node.NextOne
	}
	next.NextOne = x
	s.len++
	return nil

}

func (s *Slist) Remove() error {
	if s.len == 0 {
		return errors.New("Slist,Empty")
	}
	var next *Element = nil
	if s.len == 1 {
		s.len--
		s.node = nil
		return nil
	}
	for i := uint64(0); i <= (s.len - 2); i++ {
		next = s.node.NextOne
	}
	next.NextOne = nil
	s.len--
	return nil
}

func (s *Slist) Get(size uint64) (*Element, error) {
	if size == 0 {
		return nil, errors.New("When Slist is empty,size==0")
	}
	var next *Element
	if size == 1 {
		return s.node, nil
	}
	for i := uint64(0); i <= (size - 2); i++ {
		next = s.node.NextOne
	}
	return next, nil
}

func (s *Slist) String() string {
	var st strings.Builder
	fmt.Println(s.len)
	if s.len == 0 {
		return ""
	}
	for i := uint64(0); i < s.len; i++ {
		fmt.Println(i)
		value, err := s.Get((i + 1))
		if err != nil {
			panic(err)
		}
		st.WriteString((*value).String())
		st.WriteString("->")
	}
	return st.String()
}

func (s *Slist) Len() uint64 {
	return s.len
}
