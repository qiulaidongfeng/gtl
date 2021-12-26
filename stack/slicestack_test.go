package stack

import (
	"fmt"
	. "gtl/stack"
	"testing"
)

func Benchmark_Newslicestack(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Newslicestack()
	}
}

func Benchmark_Size(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Size()
	}
}

func Benchmark_Clear(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func Benchmark_Push(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	l := b.N
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < l; i++ {
		s.Push(5326)
	}
}

func Benchmark_Pop(b *testing.B) {
	s := Newslicestack()
	k := b.N
	var err error
	for i := 0; i < k; i++ {
		s.Push(5326)
	}
	b.ResetTimer()
	for i := 0; i < k; i++ {
		b.ReportAllocs()
		b.SetBytes(2)
		_, err = s.Pop()
		fmt.Print("\r", err)
		if err != nil {
			panic(err)
		}
	}
}

// func Benchmark_Look(b *testing.B) {
// 	s := Newslicestack()
// 	b.ReportAllocs()
// 	b.SetBytes(2)
// 	var err error
// 	for i := 0; i < b.N; i++ {
// 		s.Push(5326)
// 	}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_, err = s.Look(6)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
