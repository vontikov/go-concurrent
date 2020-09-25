package concurrent

import (
	"sync"
)

// SynchronizedSet is a safe for concurrent use Set implementation
type SynchronizedSet struct {
	mux      sync.RWMutex
	data     map[interface{}]bool
	capacity int
}

// NewSynchronizedSet returns pointer to a new SynchronizedSet instance
func NewSynchronizedSet(capacity int) Set {
	return &SynchronizedSet{
		data:     make(map[interface{}]bool, capacity),
		capacity: capacity,
	}
}

// Size implements Set.Size
func (s *SynchronizedSet) Size() int {
	s.mux.RLock()
	r := len(s.data)
	s.mux.RUnlock()
	return r
}

// Clear implements Set.Clear
func (s *SynchronizedSet) Clear() {
	s.mux.Lock()
	s.data = make(map[interface{}]bool, s.capacity)
	s.mux.Unlock()
}

// Add implements Set.Add
func (s *SynchronizedSet) Add(v interface{}) {
	s.mux.Lock()
	s.data[v] = true
	s.mux.Unlock()
}

// Contains implements Set.Contains
func (s *SynchronizedSet) Contains(v interface{}) bool {
	s.mux.RLock()
	r := s.data[v]
	s.mux.RUnlock()
	return r
}
