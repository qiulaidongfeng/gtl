package gsync

import (
	"sync"
	"testing"
)

func BenchmarkLockAndUnlock(b *testing.B) {
	lock := LMrwmutex{}
	b.Run("Parallel", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lock.Lock()
				lock.Unlock()
			}
		})
	})
	b.Run("1P", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkRLockAndRUnlock(b *testing.B) {
	lock := LMrwmutex{}
	b.Run("Parallel", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lock.RLock()
				lock.RUnlock()
			}
		})
	})
	b.Run("1P", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lock.RLock()
			lock.RUnlock()
		}
	})
}

func BenchmarkStdLockAndUnlock(b *testing.B) {
	lock := sync.RWMutex{}
	b.Run("Parallel", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lock.Lock()
				lock.Unlock()
			}
		})
	})
	b.Run("1P", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkStdRLockAndRUnlock(b *testing.B) {
	lock := sync.RWMutex{}
	b.Run("Parallel", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lock.RLock()
				lock.RUnlock()
			}
		})
	})
	b.Run("1P", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lock.RLock()
			lock.RUnlock()
		}
	})
}
