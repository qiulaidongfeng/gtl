// poolpro_test.go
package ram

import (
	. "gtl/ram"
	"testing"
)

func Benchmark_Poolpro_Len(b *testing.B) {
	b.SetBytes(1)
	n := Newpoolpro()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Len()
	}
}

func Benchmark_Poolpro_Tslen(b *testing.B) {
	b.SetBytes(1)
	n := Newpoolpro()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Tslen()
	}
}

func Benchmark_Poolpro_Newpoolpro(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Newpoolpro()
	}
}

func Benchmark_Poolpro_Put(b *testing.B) {
	b.SetBytes(1)
	n := Newpoolpro()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var k int
		n.Put(k)
	}
}

func Benchmark_Poolpro_Tsput(b *testing.B) {
	b.SetBytes(1)
	n := Newpoolpro()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var k int
		n.Tsput(k)
	}
}

func Benchmark_Poolpro_Get(b *testing.B) {
	var p = Newpoolpro()
	for i := 0; i < b.N; i++ {
		h := i
		p.Put(h)
	}
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var err error
		_, err = p.Get()
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_Poolpro_Tsget(b *testing.B) {
	var p = Newpoolpro()
	for i := 0; i < b.N; i++ {
		h := i
		p.Put(h)
	}
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var err error
		_, err = p.Tsget()
		if err != nil {
			panic(err)
		}
	}
}
