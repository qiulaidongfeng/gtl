// slist
package list

import (
	"sync"
)

type Element struct {
	value interface{}
	Next  *Element
	mutex sync.RWMutex
}

func (e *Element) Set(x interface{}) error {
	e.value = x
	return nil
}

func (e *Element) Get() (interface{}, error) {
	return e.value, nil
}

func (e *Element) Next() *Node {
	return e.Next
}

func (e *Element) Prev() *Node {
	return nil
}

type Slist struct {
}
