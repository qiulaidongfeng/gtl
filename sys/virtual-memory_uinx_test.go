package sys

import (
	"testing"
	"unsafe"
)

func TestVM(t *testing.T) {
	var p []byte = make([]byte, 64)
	v := NewVM(uintptr(unsafe.Pointer(&p[0])), uintptr(len(p)))
	err := v.Lockerr()
	if err != nil {
		t.Fatal(err)
	}
	err = v.Unlockerr()
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkVM(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p []byte = make([]byte, 64)
		v := NewVM(uintptr(unsafe.Pointer(&p[0])), uintptr(len(p)))
		err := v.Lockerr()
		if err != nil {
			b.Fatal(err)
		}
		err = v.Unlockerr()
		if err != nil {
			b.Fatal(err)
		}
	}

}
