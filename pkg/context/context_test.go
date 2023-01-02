package context

import (
	"context"
	"sync"
	"testing"
)

func TestMapContext(t *testing.T) {
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	m1 := make(map[interface{}]interface{})
	m1["a"] = "b"
	ctx2 := WithValue(ctx1, GoMap(m1))
	if ctx2.Value("a") != "b" {
		t.Fatal(ctx2.Value("a"))
	}
}

func TestSyncMapContext(t *testing.T) {
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	var m1 sync.Map
	m1.Store("a", "b")
	ctx2 := WithValue(ctx1, &m1)
	if ctx2.Value("a") != "b" {
		t.Fatal(ctx2.Value("a"))
	}

}

func BenchmarkMapContext(b *testing.B) {
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()
	m1 := make(map[interface{}]interface{})
	m1["a"] = "b"
	m2 := GoMap(m1)
	for i := 0; i < b.N; i++ {
		ctx2 := WithValue(ctx1, m2)
		if ctx2.Value("a") != "b" {
			b.Fatal(ctx2.Value("a"))
		}
	}
}

func BenchmarkSyncMapContext(b *testing.B) {
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	b.SetBytes(2)
	b.ReportAllocs()
	b.ResetTimer()
	var m1 sync.Map
	m1.Store("a", "b")
	for i := 0; i < b.N; i++ {
		ctx2 := WithValue(ctx1, &m1)
		if ctx2.Value("a") != "b" {
			b.Fatal(ctx2.Value("a"))
		}
	}
}
