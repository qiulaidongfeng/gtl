package stack

import (
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

func Benchmark_Tssize(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Tssize()
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

func Benchmark_Tsclear(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Tsclear()
	}
}

func Benchmark_Push(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(5326)
	}
}

func Benchmark_Tspush(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Tspush(5326)
	}
}

func Benchmark_Pop(b *testing.B) {
	s := Newslicestack()
	var err error
	b.ReportAllocs()
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		s.Push(90)
		_, err = s.Pop()
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_Tspop(b *testing.B) {
	s := Newslicestack()
	var err error
	b.ReportAllocs()
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		s.Push(77)
		_, err = s.Tspop()
		if err != nil {
			panic(err)
		}
	}
}

// func Benchmark_Look(b *testing.B) {
// 	s := Newslicestack()
// 	b.ReportAllocs()
// 	b.SetBytes(2)
// 	b.ResetTimer()
// 	var err error
// 	b.StopTimer()
// 	for i := 0; i < b.N; i++ {
// 		s.Push(5326)
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		_, err = s.Look(6)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
