// silcestack-test_test
package stack

import (
	"testing"
)

var (
	err  error
	x    interface{}
	size int = 100000
)

func Test_Silce_Size(t *testing.T) {
	s := Newslicestack()
	for i := 0; i < size; i++ {
		err = s.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	stacksize := s.Size()
	if stacksize != uint64(size) {
		t.Fatal("Size不正确")
	}
}

func Test_Silce_TsSize(t *testing.T) {
	s := Newslicestack()
	for i := 0; i < size; i++ {
		err = s.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	stacksize := s.TsSize()
	if stacksize != uint64(size) {
		t.Fatal("Size不正确")
	}
}

func Test_Silce_Clear(t *testing.T) {
	s := Newslicestack()
	for i := 0; i < size; i++ {
		err = s.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	err := s.Clear()
	if err != nil {
		t.Fatal(err)
	}
	stacksize := s.Size()
	if stacksize != 0 {
		t.Fatal("Size不正确")
	}
}

func Test_Silce_TsClear(t *testing.T) {
	s := Newslicestack()
	for i := 0; i < size; i++ {
		err = s.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	err := s.TsClear()
	if err != nil {
		t.Fatal(err)
	}
	stacksize := s.Size()
	if stacksize != 0 {
		t.Fatal("Size不正确")
	}
}

func Test_Silce_PushAndPop(t *testing.T) {
	s := Newslicestack()
	for i := 0; i < size; i++ {
		err = s.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := size - 1; i > 0; i-- {
		x, err = s.Pop()
		if x.(int) != i {
			t.Fatal("出栈的数据错误")
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}

func Test_Silce_TsPushAndTsPop(t *testing.T) {
	s := Newslicestack()
	for i := 0; i < size; i++ {
		err = s.TsPush(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := size - 1; i > 0; i-- {
		x, err = s.TsPop()
		if x.(int) != i {
			t.Fatal("出栈的数据错误")
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}
