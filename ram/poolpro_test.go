// poolpro_test.go
package ram

import (
	. "gtl/ram"
	"testing"
)

var (
	p = Newpoolpro()
)

func oneinit(i int) {
	for io := 0; io < i; i++ {
		k := 90
		p.Put(k)
	}
}

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
	oneinit(2)
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < 2; i++ {
		var err error
		_, err = p.Get()
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_Poolpro_Tsget(b *testing.B) {
	oneinit(2)
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < 2; i++ {
		var err error
		_, err = p.Tsget()
		if err != nil {
			panic(err)
		}
	}
}
