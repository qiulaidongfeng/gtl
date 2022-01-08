// slist_test.go
package list

import (
	"testing"

	. "gtl/list"
)

func Benchmark_Newslist(b *testing.B) {
	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSlist()
	}
}

func Benchmark_Lnsert(b *testing.B) {
	s := NewSlist()
	e := Element{}
	e.Set(80)
	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := s.Lnsert(&e)
		if err != nil {
			panic(err)
		}
	}
}
