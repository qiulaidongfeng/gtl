package cextend

import (
	"testing"
	"unsafe"
)

var (
	ptr unsafe.Pointer
)

func TestMalloc(t *testing.T) {
	ptr = Malloc(8)
	if ptr == nil {
		t.Fatal("没有分配内存！")
	}
}

func TestFree(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Fatal(err)
		}
	}()
	Free(ptr)
}

func TestMemcpy(t *testing.T) {
	const size uint = 1000
	var a, b [size]int8
	for i := uint(0); i < size; i++ {
		a[i] = int8(i / (1 << 8))
	}
	Memcpy(unsafe.Pointer(&b[0]), unsafe.Pointer(&a[0]), size)
	for i := uint(0); i < size; i++ {
		if !(b[i] == a[i]) {
			t.Fatal("复制出现错误！")
		}
	}
}

func TestCalloc(t *testing.T) {
	ptr = Calloc(1, 8)
	if ptr == nil {
		t.Fatal("没有分配内存！")
	}
}

func BenchmarkMallocAndFree(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		ptr = Malloc(8)
		Free(ptr)
	}
}

func BenchmarkCallocAndFree(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		ptr = Calloc(1, 8)
		Free(ptr)
	}
}

func BenchmarkMemcpy_100bit(b *testing.B) {
	b.StopTimer()
	const size uint = 100
	var a, c [size]int8
	for i := uint(0); i < size; i++ {
		a[i] = int8(i / (1 << 8))
	}
	b.ReportAllocs()
	b.SetBytes(1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Memcpy(unsafe.Pointer(&c[0]), unsafe.Pointer(&a[0]), size)
	}
	for i := uint(0); i < size; i++ {
		if !(c[i] == a[i]) {
			b.Fatal("复制出现错误！")
		}
	}
}

func BenchmarkMemcpy_1000bit(b *testing.B) {
	b.StopTimer()
	const size uint = 1000
	var a, c [size]int8
	for i := uint(0); i < size; i++ {
		a[i] = int8(i / (1 << 8))
	}
	b.ReportAllocs()
	b.SetBytes(1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Memcpy(unsafe.Pointer(&c[0]), unsafe.Pointer(&a[0]), size)
	}
	for i := uint(0); i < size; i++ {
		if !(c[i] == a[i]) {
			b.Fatal("复制出现错误！")
		}
	}
}

func BenchmarkMemcpy_1000000bit(b *testing.B) {
	b.StopTimer()
	const size uint = 1000000
	var a, c [size]int8
	for i := uint(0); i < size; i++ {
		a[i] = int8(i / (1 << 8))
	}
	b.ReportAllocs()
	b.SetBytes(1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Memcpy(unsafe.Pointer(&c[0]), unsafe.Pointer(&a[0]), size)
	}
	for i := uint(0); i < size; i++ {
		if !(c[i] == a[i]) {
			b.Fatal("复制出现错误！")
		}
	}
}
