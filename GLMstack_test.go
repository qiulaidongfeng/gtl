// GLMstack_test
package stack

import (
	"testing"
	"unsafe"

	"gtl/stack"
)

var s *stack.GLMstack = stack.NewGLMstack()

func Benchmark_NewGLMstack(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = stack.NewGLMstack()
	}
}

func Benchmark_Size(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Size()
	}
}

func Benchmark_Tssize(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Tssize()
	}
}

func Benchmark_Clear(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func Benchmark_Tsclear(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Tsclear()
	}
}

func Benchmark_TsLoadpp(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.TsLoadpp()
	}
}

func Benchmark_TsLoadpopn(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.TsLoadpopn()
	}
}

func Benchmark_TsLoadpushn(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.TsLoadpushn()
	}
}

func Benchmark_Pushptr(b *testing.B) {
	var s1 *stack.GLMstack = stack.NewGLMstack()
	p := new(int)
	up := unsafe.Pointer(p)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.Pushptr(up, 1)
	}
}

func Benchmark_TsPushptr(b *testing.B) {
	var s1 *stack.GLMstack = stack.NewGLMstack()
	p := new(int)
	up := unsafe.Pointer(p)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.TsPushptr(up, 1)
	}
}

func Benchmark_Pushint8(b *testing.B) {
	var s1 *stack.GLMstack = stack.NewGLMstack()
	p := int8(9)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.Pushint8(p)
	}
}

func Benchmark_TsPushint8(b *testing.B) {
	var s1 *stack.GLMstack = stack.NewGLMstack()
	p := int8(9)
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1.TsPushint8(p)
	}
}
