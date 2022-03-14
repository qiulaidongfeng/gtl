package stack

import (
	"testing"
)

func Benchmark_Slice_NewSlicestack(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Newslicestack()
	}
}

func Benchmark_Silce_Size(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Size()
	}
}

func Benchmark_Slice_TsSize(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.TsSize()
	}
}

func Benchmark_Slice_Clear(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func Benchmark_Slice_TsClear(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.TsClear()
	}
}

func Benchmark_Slice_Push(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(5326)
	}
}

func Benchmark_Slice_PushAndPop(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(5326)
		_, err = s.Pop()
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_Slice_TsPushAndTsPop(b *testing.B) {
	b.SetBytes(2)
	s := Newslicestack()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.TsPush(5326)
		_, err = s.TsPop()
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
