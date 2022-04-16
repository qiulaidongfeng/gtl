//go:build go1.18
// +build go1.18

package bits

import (
	"math/rand"
	"testing"
)

type int16bits_setbit1 struct {
	in    int16
	newin int16
}

func newint16bits_setbit1(in int16) *int16bits_setbit1 {
	b := &int16bits_setbit1{
		in: in,
	}
	return b
}

func (b *int16bits_setbit1) Setbit1(index int16) int16 {
	b.newin = Setbit1(b.in, index)
	return b.newin
}

func (b *int16bits_setbit1) Check(indexr1 int16) bool {
	for index := int16(0); index < 16; index++ {
		inbit := Getbit(b.in, index)
		newbit := Getbit(b.newin, index)
		if inbit == newbit {
			continue
		}
		if index == indexr1 {
			continue
		}
		return false
	}
	return true
}

func FuzzSetbit1_Int16(f *testing.F) {
	for _, value := range []int16{int16(rand.Int31n(1 << 16)), int16(rand.Int31n(1 << 16)), 0} {
		f.Add(value)
	}
	f.Fuzz(func(t *testing.T, in int16) {
		index := int16(rand.Int31n(16))
		bstruct := newint16bits_setbit1(in)
		bitold := bstruct.Setbit1(index)
		ok := bstruct.Check(index)
		if !ok {
			t.Fatalf("%d位\n%b\n%b\n%s", index, in, bitold, "比特位值获取不正确！")
		}
	})
}

type int16bits_setbit0 struct {
	in    int16
	newin int16
}

func newint16bits_setbit0(in int16) *int16bits_setbit0 {
	b := &int16bits_setbit0{
		in: in,
	}
	return b
}

func (b *int16bits_setbit0) Setbit0(index int16) int16 {
	b.newin = Setbit0(b.in, index)
	return b.newin
}

func (b *int16bits_setbit0) Check(indexr0 int16) bool {
	for index := int16(0); index < 16; index++ {
		inbit := Getbit(b.in, index)
		newbit := Getbit(b.newin, index)
		if inbit == newbit {
			continue
		}
		if index == indexr0 {
			continue
		}
		return false
	}
	return true
}

func FuzzSetbit0_Int16(f *testing.F) {
	for _, value := range []int16{int16(rand.Int31n(1 << 16)), int16(rand.Int31n(1 << 16)), 0} {
		f.Add(value)
	}
	f.Fuzz(func(t *testing.T, in int16) {
		index := int16(rand.Int31n(16))
		bstruct := newint16bits_setbit0(in)
		bitold := bstruct.Setbit0(index)
		ok := bstruct.Check(index)
		if !ok {
			t.Fatalf("%d位\n%b\n%b\n%s", index, in, bitold, "比特位值获取不正确！")
		}
	})
}

func TestGetbit_Int64(t *testing.T) {
	a := make([]int64, 64)
	for i := int64(1); i < 64; i++ {
		a[i] = 1 << i
	}
	for i := int64(1); i < 64; i++ {
		bit := Getbit(a[i], i)
		if bit != 1 {
			t.Fatal(i, "比特位值获取不正确！")
		}
	}
}

func TestSetbit1_Int64(t *testing.T) {
	a := make([]int64, 64)
	for i := int64(1); i < 64; i++ {
		a[i] = rand.Int63()
	}
	for i := int64(1); i < 64; i++ {
		bit := Setbit1(a[i], i)
		bitnew := Getbit(bit, i)
		if bitnew != 1 {
			t.Fatalf("%d位\n%b\n%b\n%s", i, a[i], bit, "比特位值设置不正确！")
		}
	}
}

func TestSetbit0_Int64(t *testing.T) {
	a := make([]int64, 64)
	for i := int64(1); i < 64; i++ {
		a[i] = rand.Int63()
	}
	for i := int64(1); i < 64; i++ {
		bit := Setbit0(a[i], i)
		bitnew := Getbit(bit, i)
		if bitnew != 0 {
			t.Fatalf("%d位\n%b\n%b\n%s", i, a[i], bit, "比特位值设置不正确！")
		}
	}
}

func BenchmarkGetbit_Int64(b *testing.B) {
	b.SetBytes(99)
	b.ReportAllocs()
	a := 1 << 32
	for i := int64(1); i < int64(b.N); i++ {
		_ = Getbit(a, 32)
	}
}
func BenchmarkSetbit1_Int64(b *testing.B) {
	b.SetBytes(99)
	b.ReportAllocs()
	a := rand.Int63()
	for i := int64(1); i < int64(b.N); i++ {
		_ = Setbit1(a, 32)
	}
}
func BenchmarkSetbit0_Int64(b *testing.B) {
	b.SetBytes(99)
	b.ReportAllocs()
	a := rand.Int63()
	for i := int64(1); i < int64(b.N); i++ {
		_ = Setbit0(a, 32)
	}
}
