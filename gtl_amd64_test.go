package gtl

import (
	"math/rand"
	"testing"
	"unsafe"
)

const (
	Men1kb   = 1024
	Men100kb = Men1kb * 100
	Men1mb   = Men1kb * 1024
)

func TestCopy(t *testing.T) {
	var (
		s1 = [10]int64{32, 547, 973, 223, 76, 1, 76, 80, 97, 43}
		d1 = [10]int64{0, 0, 2, 0, 0, 0, 0, 0, 0, 0}
		s2 = [10]byte{71, 43, 13, 99, 221, 209, 21, 22, 78, 103}
		d2 = [10]byte{0, 2, 0, 0, 0, 0, 0, 0, 0, 0}
		s3 = [2]byte{71, 43}
		d3 = [2]byte{0, 1}
	)
	type args struct {
		dest unsafe.Pointer
		src  unsafe.Pointer
		n    uint
		name string
	}
	setargs := func(arg *args, dest, src unsafe.Pointer, n uint, name string) {
		arg.dest = dest
		arg.src = src
		arg.n = n
		arg.name = name
		return
	}
	var tests = make([]args, 3, 3)
	setargs(&tests[0], unsafe.Pointer(&d1[0]), unsafe.Pointer(&s1[0]), (uint(len(d1)) * uint(8)), "80字节")
	setargs(&tests[1], unsafe.Pointer(&d3[0]), unsafe.Pointer(&s3[0]), uint(len(d3)), "2字节")
	setargs(&tests[2], unsafe.Pointer(&d2[0]), unsafe.Pointer(&s2[0]), uint(len(d2)), "10字节")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Copy(test.dest, test.src, test.n)
			var di, si []byte
			di = mtob(uintptr(test.dest), uintptr(test.n))
			si = mtob(uintptr(test.src), uintptr(test.n))
			for i := 0; i < len(si); i++ {
				sicmp := si[i]
				dicmp := di[i]
				if sicmp != dicmp {
					t.Fatal(si, di)
				}
			}
		})
	}
}

func TestCopy_Movups(t *testing.T) {
	var (
		s1 = [10]int64{32, 547, 973, 223, 76, 1, 76, 80, 97, 43}
		d1 = [10]int64{0, 0, 2, 0, 0, 0, 0, 0, 0, 0}
		s2 = [10]byte{71, 43, 13, 99, 221, 209, 21, 22, 78, 103}
		d2 = [10]byte{0, 2, 0, 0, 0, 0, 0, 0, 0, 0}
		s3 = [2]byte{71, 43}
		d3 = [2]byte{0, 1}
	)
	type args struct {
		dest unsafe.Pointer
		src  unsafe.Pointer
		n    uint
		name string
	}
	setargs := func(arg *args, dest, src unsafe.Pointer, n uint, name string) {
		arg.dest = dest
		arg.src = src
		arg.n = n
		arg.name = name
		return
	}
	var tests = make([]args, 3, 3)
	setargs(&tests[0], unsafe.Pointer(&d1[0]), unsafe.Pointer(&s1[0]), (uint(len(d1)) * uint(8)), "80字节")
	setargs(&tests[1], unsafe.Pointer(&d3[0]), unsafe.Pointer(&s3[0]), uint(len(d3)), "2字节")
	setargs(&tests[2], unsafe.Pointer(&d2[0]), unsafe.Pointer(&s2[0]), uint(len(d2)), "10字节")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Copy_Movups(test.dest, test.src, test.n)
			var di, si []byte
			di = mtob(uintptr(test.dest), uintptr(test.n))
			si = mtob(uintptr(test.src), uintptr(test.n))
			for i := 0; i < len(si); i++ {
				sicmp := si[i]
				dicmp := di[i]
				if sicmp != dicmp {
					t.Fatal(si, di)
				}
			}
		})
	}
}

func TestMemmove(t *testing.T) {
	var (
		s1 = [10]int64{32, 547, 973, 223, 76, 1, 76, 80, 97, 43}
		d1 = [10]int64{0, 0, 2, 0, 0, 0, 0, 0, 0, 0}
		s2 = [10]byte{71, 43, 13, 99, 221, 209, 21, 22, 78, 103}
		d2 = [10]byte{0, 2, 0, 0, 0, 0, 0, 0, 0, 0}
		s3 = [2]byte{71, 43}
		d3 = [2]byte{0, 1}
	)
	type args struct {
		dest unsafe.Pointer
		src  unsafe.Pointer
		n    uint
		name string
	}
	setargs := func(arg *args, dest, src unsafe.Pointer, n uint, name string) {
		arg.dest = dest
		arg.src = src
		arg.n = n
		arg.name = name
		return
	}
	var tests = make([]args, 3, 3)
	setargs(&tests[0], unsafe.Pointer(&d1[0]), unsafe.Pointer(&s1[0]), (uint(len(d1)) * uint(8)), "80字节")
	setargs(&tests[1], unsafe.Pointer(&d3[0]), unsafe.Pointer(&s3[0]), uint(len(d3)), "2字节")
	setargs(&tests[2], unsafe.Pointer(&d2[0]), unsafe.Pointer(&s2[0]), uint(len(d2)), "10字节")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Memmove(test.dest, test.src, test.n)
			var di, si []byte
			di = mtob(uintptr(test.dest), uintptr(test.n))
			si = mtob(uintptr(test.src), uintptr(test.n))
			for i := 0; i < len(si); i++ {
				sicmp := si[i]
				dicmp := di[i]
				if sicmp != dicmp {
					t.Fatal(si, di)
				}
			}
		})
	}
}

//go:uintptrescapes
func mtob(addr, length uintptr) []byte {
	var a []byte
	p := (*[3]uintptr)(unsafe.Pointer(&a))
	p[0] = addr
	p[1] = length
	p[2] = length
	return a
}

func FuzzCopy(f *testing.F) {
	var s1 = []byte("264872")
	var d1 = []byte("264872")
	f.Add(s1, d1)
	f.Fuzz(func(t *testing.T, si, di []byte) {
		if len(si) > len(di) || len(si) == 0 || len(di) == 0 {
			return
		}
		Copy(unsafe.Pointer(&di[0]), unsafe.Pointer(&si[0]), uint(len(si)))
		for i := 0; i < (len(si)); i++ {
			sicmp := si[i]
			dicmp := di[i]
			if sicmp != dicmp {
				t.Fatal(si, di)
			}
		}
	})
}

func BenchmarkCopy_Movups_1mb(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	var (
		si [Men1mb]int8
		di [Men1mb]int8
	)
	for i := 0; i < Men1mb; i++ {
		si[i] = int8(rand.Int63())
		di[i] = 0
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy_Movups(unsafe.Pointer(&di[0]), unsafe.Pointer(&si[0]), Men1mb)
	}
}

func BenchmarkCopy_1mb(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	var (
		si [Men1mb]int8
		di [Men1mb]int8
	)
	for i := 0; i < Men1mb; i++ {
		si[i] = int8(rand.Int63())
		di[i] = 0
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(unsafe.Pointer(&di[0]), unsafe.Pointer(&si[0]), Men1mb)
	}
}

func BenchmarkMemmove_1mb(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	var (
		si [Men1mb]int8
		di [Men1mb]int8
	)
	for i := 0; i < Men1mb; i++ {
		si[i] = int8(rand.Int63())
		di[i] = 0
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Memmove(unsafe.Pointer(&di[0]), unsafe.Pointer(&si[0]), Men1mb)
	}
}
