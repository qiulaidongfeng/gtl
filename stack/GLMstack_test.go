// GLMstack_test
package stack

import (
	"testing"
	"unsafe"
)

var s *GLMstack = NewGLMstack()

func Benchmark_GLM_NewGLMstack(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = NewGLMstack()
	}
}

func Benchmark_GLM_Size(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Size()
	}
}

func Benchmark_GLM_Tssize(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Tssize()
	}
}

func Benchmark_GLM_Clear(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func Benchmark_GLM_Tsclear(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Tsclear()
	}
}

func Benchmark_GLM_TsLoadpp(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.TsLoadpp()
	}
}

func Benchmark_GLM_TsLoadpopn(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.TsLoadpopn()
	}
}

func Benchmark_GLM_TsLoadpushn(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.TsLoadpushn()
	}
}

func Benchmark_GLM_Pushptr(b *testing.B) {
	var s1 *GLMstack = NewGLMstack()
	p := new(int)
	up := unsafe.Pointer(p)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.Pushptr(up, 1)
	}
}

func Benchmark_GLM_TsPushptr(b *testing.B) {
	var s1 *GLMstack = NewGLMstack()
	p := new(int)
	up := unsafe.Pointer(p)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.TsPushptr(up, 1)
	}
}

func Benchmark_GLM_Pushint8(b *testing.B) {
	var s1 *GLMstack = NewGLMstack()
	p := int8(9)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.Pushint8(p)
	}
}

func Benchmark_GLM_TsPushint8(b *testing.B) {
	var s1 *GLMstack = NewGLMstack()
	p := int8(9)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.TsPushint8(p)
	}
}
