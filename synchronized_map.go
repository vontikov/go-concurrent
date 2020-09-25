package concurrent

import (
	"sync"
)

// SynchronizedMap is a safe for concurrent use Map implementation
type SynchronizedMap struct {
	mux  sync.RWMutex
	data map[interface{}]interface{}
}

// NewSynchronizedMap returns pointer to a new SynchronizedMap instance
func NewSynchronizedMap(capacity int) Map {
	return &SynchronizedMap{
		data: make(map[interface{}]interface{}, capacity),
	}
}

// Size implements Map.Size
func (m *SynchronizedMap) Size() int {
	m.mux.RLock()
	r := len(m.data)
	m.mux.RUnlock()
	return r
}

// Clear implements Map.Clear
func (m *SynchronizedMap) Clear() {
	m.mux.Lock()
	m.data = make(map[interface{}]interface{})
	m.mux.Unlock()
}

// Put implements Map.Put
func (m *SynchronizedMap) Put(k interface{}, v interface{}) interface{} {
	m.mux.Lock()
	o, _ := m.data[k]
	m.data[k] = v
	m.mux.Unlock()
	return o
}

// Get implements Map.Get
func (m *SynchronizedMap) Get(k interface{}) interface{} {
	m.mux.RLock()
	r := m.data[k]
	m.mux.RUnlock()
	return r
}

// PutIfAbsent implements Map.PutIfAbsent
func (m *SynchronizedMap) PutIfAbsent(k interface{}, v interface{}) bool {
	m.mux.Lock()
	if _, ok := m.data[k]; ok {
		m.mux.Unlock()
		return false
	}
	m.data[k] = v
	m.mux.Unlock()
	return true
}

// Range implements Map.Range
func (m *SynchronizedMap) Range(f func(k, v interface{}) bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			return
		}
	}
}
