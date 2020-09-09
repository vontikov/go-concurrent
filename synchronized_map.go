package concurrent

import (
	"sync"
)

// SynchronizedMap is a safe for concurrent use Map implementation
type SynchronizedMap struct {
	sync.RWMutex
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
	m.RLock()
	r := len(m.data)
	m.RUnlock()
	return r
}

// Clear implements Map.Clear
func (m *SynchronizedMap) Clear() {
	m.Lock()
	m.data = make(map[interface{}]interface{})
	m.Unlock()
}

// Put implements Map.Put
func (m *SynchronizedMap) Put(k interface{}, v interface{}) {
	m.Lock()
	m.data[k] = v
	m.Unlock()
}

// PutIfAbsent implements Map.PutIfAbsent
func (m *SynchronizedMap) PutIfAbsent(k interface{}, v interface{}) bool {
	m.Lock()
	if _, ok := m.data[k]; ok {
		m.Unlock()
		return false
	}
	m.data[k] = v
	m.Unlock()
	return true
}

// Get implements Map.Get
func (m *SynchronizedMap) Get(k interface{}) interface{} {
	m.RLock()
	r := m.data[k]
	m.RUnlock()
	return r
}
