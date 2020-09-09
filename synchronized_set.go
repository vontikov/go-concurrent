package concurrent

import (
	"sync"
)

// SynchronizedSet is a safe for concurrent use Set implementation
type SynchronizedSet struct {
	sync.RWMutex
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
	s.RLock()
	r := len(s.data)
	s.RUnlock()
	return r
}

// Clear implements Set.Clear
func (s *SynchronizedSet) Clear() {
	s.Lock()
	s.data = make(map[interface{}]bool, s.capacity)
	s.Unlock()
}

// Add implements Set.Add
func (s *SynchronizedSet) Add(v interface{}) {
	s.Lock()
	s.data[v] = true
	s.Unlock()
}

// Contains implements Set.Contains
func (s *SynchronizedSet) Contains(v interface{}) bool {
	s.RLock()
	r := s.data[v]
	s.RUnlock()
	return r
}
