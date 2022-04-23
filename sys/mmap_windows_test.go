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
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := mmap.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
}

func BenchmarkNewMmap(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()
	defer func() {
		err := recover()
		if err != nil {
			b.Fatal(err)
		}
	}()
	for i := 0; i < b.N; i++ {
		mmap, err := NewMmap("./test/mmap_test.txt", 0)
		if err != nil {
			b.Fatal(err)
		}
		defer func() {
			err := mmap.Close()
			if err != nil {
				b.Fatal(err)
			}
		}()
	}

}
