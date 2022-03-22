package sys

import (
	"testing"
)

func TestNewMmap(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Fatal(err)
		}
	}()
	mmap, err := NewMmap("./test/mmap_test.txt", 0)
	defer func() {
		err := mmap.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
}
