package concurrent

import (
	"sync"
)

// SynchronizedList is a safe for concurrent use List implementation
type SynchronizedList struct {
	sync.RWMutex
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
	l.RLock()
	r := len(l.data)
	l.RUnlock()
	return r
}

// Clear implements List.Clear
func (l *SynchronizedList) Clear() {
	l.Lock()
	l.data = make([]interface{}, 0, l.capacity)
	l.Unlock()
}

// Add implements List.Clear
func (l *SynchronizedList) Add(v interface{}) {
	l.Lock()
	l.data = append(l.data, v)
	l.Unlock()
}

// Get implements List.Clear
func (l *SynchronizedList) Get(i int) interface{} {
	l.RLock()
	v := l.data[i]
	l.RUnlock()
	return v
}
