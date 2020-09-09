package concurrent

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkInsertSyncMap(b *testing.B) {
	var m sync.Map
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(strconv.Itoa(i), "value")
	}
}

func BenchmarkInsertSynchronizedMap(b *testing.B) {
	m := NewSynchronizedMap(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Put(strconv.Itoa(i), "value")
	}
}

func BenchmarkInsertSynchronizedMapPreAllocated(b *testing.B) {
	m := NewSynchronizedMap(10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Put(strconv.Itoa(i), "value")
	}
}
