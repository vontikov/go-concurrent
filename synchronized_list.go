package concurrent

import (
	"sync"
)

// SynchronizedList is a safe for concurrent use List implementation
type SynchronizedList struct {
	mux      sync.RWMutex
	data     []interface{}
	capacity int
}

// NewSynchronizedList returns pointer to a new SynchronizedList instance
func NewSynchronizedList(capacity int) List {
	return &SynchronizedList{
		data:     make([]interface{}, 0, capacity),
		capacity: capacity,
	}
}

// Size implements List.Size
func (l *SynchronizedList) Size() int {
	l.mux.RLock()
	r := len(l.data)
	l.mux.RUnlock()
	return r
}

// Clear implements List.Clear
func (l *SynchronizedList) Clear() {
	l.mux.Lock()
	l.data = make([]interface{}, 0, l.capacity)
	l.mux.Unlock()
}

// Add implements List.Clear
func (l *SynchronizedList) Add(v interface{}) {
	l.mux.Lock()
	l.data = append(l.data, v)
	l.mux.Unlock()
}

// Get implements List.Clear
func (l *SynchronizedList) Get(i int) interface{} {
	l.mux.RLock()
	v := l.data[i]
	l.mux.RUnlock()
	return v
}
