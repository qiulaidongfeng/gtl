//go:build go1.18
// +build go1.18

package bits

import (
	"math/rand"
	"testing"
)

type int64bits_setbit1 struct {
	in    int64
	newin int64
}

func newint64bits_setbit1(in int64) int64bits_setbit1 {
	b := int64bits_setbit1{
		in: in,
	}
	return b
}

func (b int64bits_setbit1) Setbit1(index int64) int64 {
	b.newin = Setbit1(b.in, int64(index))
	return b.newin
}

func (b int64bits_setbit1) Check(indexr1 int64) (ok bool) {
	for index := int64(0); index < 64; index++ {
		inbit := Getbit(b.in, index)
		newbit := Getbit(b.newin, index)
		if inbit == newbit {
			continue
		}
		if index == indexr1 {
			continue
		}
		ok = false
	}
	return true
}

func FuzzGetbit_Int64(f *testing.F) {
	for _, value := range []int64{rand.Int63(), rand.Int63(), rand.Int63()} {
		f.Add(value)
	}
	f.Fuzz(func(t *testing.T, in int64) {
		index := rand.Int63n(63)
		bstruct := newint64bits_setbit1(in)
		bitold := bstruct.Setbit1(index)
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
	a := int64(1)
	a = a << 32
	for i := int64(1); i < int64(b.N); i++ {
		bit := Getbit(a, 32)
		if bit != 1 {
			b.Fatal(i, "比特位值设置不正确！")
		}
	}
}
