// poolpro_test.go
package ram

import (
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

func Benchmark_Poolpro_Len(b testing.B) {
	b.SetBytes()
	n := Newpoolpro()
	b.ReportMetric()
	for i := 0; i < b.N; i++ {
		n.Len()
	}
}

func Benchmark_Poolpro_Tslen(b testing.B) {
	b.SetBytes()
	n := Newpoolpro()
	b.ReportMetric()
	for i := 0; i < b.N; i++ {
		n.Tslen()
	}
}

func Benchmark_Poolpro_Newpoolpro(b testing.B) {
	b.SetBytes()
	b.ReportMetric()
	for i := 0; i < b.N; i++ {
		l := Newpoolpro()
	}
}

func Benchmark_Poolpro_Len(b testing.B) {
	b.SetBytes()
	n := Newpoolpro()
	b.ReportMetric()
	for i := 0; i < b.N; i++ {
		n.Len()
	}
}

func Benchmark_Poolpro_Put(b testing.B) {
	b.SetBytes()
	n := Newpoolpro()
	b.ReportMetric()
	for i := 0; i < b.N; i++ {
		var k int
		n.put(k)
	}
}

func Benchmark_Poolpro_Tsput(b testing.B) {
	b.SetBytes()
	n := Newpoolpro()
	b.ReportMetric()
	for i := 0; i < b.N; i++ {
		var k int
		n.Tsput(k)
	}
}

func Benchmark_Poolpro_Get(b testing.B) {
	oneinit(100000)
	b.SetBytes()
	b.ReportMetric()
	for i := 0; i < 100000; i++ {
		face, err := p.Get()
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_Poolpro_Get(b testing.B) {
	oneinit(100000)
	b.SetBytes()
	b.ReportMetric()
	for i := 0; i < 100000; i++ {
		p.Tsget()
		face, err := p.Get()
		if err != nil {
			panic(err)
		}
	}
}
