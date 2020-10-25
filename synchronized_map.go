package concurrent

import (
	"sync"
)

// SynchronizedMap is a safe for concurrent use Map implementation.
type SynchronizedMap struct {
	sync.RWMutex
	data map[interface{}]interface{}
}

// NewSynchronizedMap returns pointer to a new SynchronizedMap instance.
func NewSynchronizedMap(capacity int) Map {
	return &SynchronizedMap{
		data: make(map[interface{}]interface{}, capacity),
	}
}

// Size implements Map.Size.
func (m *SynchronizedMap) Size() int {
	m.RLock()
	r := len(m.data)
	m.RUnlock()
	return r
}

// Clear implements Map.Clear.
func (m *SynchronizedMap) Clear() {
	m.Lock()
	m.data = make(map[interface{}]interface{})
	m.Unlock()
}

// Put implements Map.Put.
func (m *SynchronizedMap) Put(k interface{}, v interface{}) interface{} {
	m.Lock()
	o, _ := m.data[k]
	m.data[k] = v
	m.Unlock()
	return o
}

// PutIfAbsent implements Map.PutIfAbsent.
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

// ComputeIfAbsent implements Map.ComputeIfAbsent.
func (m *SynchronizedMap) ComputeIfAbsent(k interface{}, f func() interface{}) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()
	if v, ok := m.data[k]; ok {
		return v, false
	}

	v := f()
	if v == nil {
		return nil, false
	}
	m.data[k] = v
	return v, true
}

// Contains implements Map.Contains.
func (m *SynchronizedMap) Contains(k interface{}) bool {
	m.RLock()
	_, ok := m.data[k]
	m.RUnlock()
	return ok
}

// Get implements Map.Get.
func (m *SynchronizedMap) Get(k interface{}) interface{} {
	m.RLock()
	r := m.data[k]
	m.RUnlock()
	return r
}

// Range implements Map.Range.
func (m *SynchronizedMap) Range(f func(k, v interface{}) bool) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			return
		}
	}
}

// Remove implements Map.Remove.
func (m *SynchronizedMap) Remove(k interface{}) {
	m.Lock()
	delete(m.data, k)
	m.Unlock()
}

// Keys implements Map.Keys.
func (m *SynchronizedMap) Keys() []interface{} {
	m.RLock()
	defer m.RUnlock()
	sz := len(m.data)
	r := make([]interface{}, 0, sz)
	for k := range m.data {
		r = append(r, k)
	}
	return r
}
