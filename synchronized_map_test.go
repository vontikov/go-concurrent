package concurrent

import (
	"testing"

	"sort"
	"sync"

	"github.com/stretchr/testify/assert"
)

func TestSynchronizedMapEmpty(t *testing.T) {
	assert.Equal(t, 0, NewSynchronizedMap(10).Size(), "Empty map size must be equal to 0")
}

func TestSynchronizedMapPutGet(t *testing.T) {
	const n = 100
	const l = n * n << 2

	m := NewSynchronizedMap(l)

	var wg sync.WaitGroup

	// populate the map
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				assert.Nil(t, m.Put(k, v))
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()

	// make sure the old values are overridden
	base = 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				assert.NotNil(t, m.Put(k, v))
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()

	assert.Equal(t, n*n, m.Size(), "All items should be added")

	s := make([]int, 0, m.Size())
	for i := 0; i < n*n; i++ {
		s = append(s, m.Get(i).(int))
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	for i := 0; i < n*n; i++ {
		assert.Equal(t, l+i, s[i], "Items should be unique")
	}
}

func TestSynchronizedMapClear(t *testing.T) {
	const n = 100

	m := NewSynchronizedMap(n)
	assert.Equal(t, 0, m.Size(), "Should be empty")

	for i := 0; i < n; i++ {
		m.Put(i, true)
	}
	assert.Equal(t, n, m.Size(), "Should not be empty")

	m.Clear()
	assert.Equal(t, 0, m.Size(), "Should be empty")
}

func TestSynchronizedMapPutIfAbsent(t *testing.T) {
	const n = 100
	const l = n * n << 2

	m := NewSynchronizedMap(l)

	// first pass
	var wg sync.WaitGroup
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				r := m.PutIfAbsent(k, v)
				assert.True(t, r, "Should add the pair")
			}
			wg.Done()
		}(base)
		base += n
	}

	wg.Wait()
	assert.Equal(t, n*n, m.Size(), "All items should be added")

	s := make([]int, 0, m.Size())
	for i := 0; i < n*n; i++ {
		s = append(s, m.Get(i).(int))
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	for i := 0; i < n*n; i++ {
		assert.Equal(t, l+i, s[i], "Items should be unique")
	}

	// second pass
	base = 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				r := m.PutIfAbsent(k, v)
				assert.False(t, r, "Should not add the pair")
			}
			wg.Done()
		}(base)
		base += n
	}

	wg.Wait()
	assert.Equal(t, n*n, m.Size(), "Items should not be added")
}

func TestSynchronizedMapRange(t *testing.T) {
	const n = 100
	const l = n * n << 2

	m := NewSynchronizedMap(l)

	var wg sync.WaitGroup
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				m.Put(k, v)
			}
			wg.Done()
		}(base)
		base += n
	}

	wg.Wait()
	assert.Equal(t, n*n, m.Size(), "All items should be added")

	var keys, values []int
	m.Range(func(k, v interface{}) bool {
		keys = append(keys, k.(int))
		values = append(values, v.(int))
		return true
	})

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

	for i := 0; i < n*n; i++ {
		assert.Equal(t, i, keys[i], "Keys should be unique")
		assert.Equal(t, l+i, values[i], "Values should be unique")
	}
}

func TestSynchronizedMapRemove(t *testing.T) {
	const n = 100
	const l = n * n << 2

	m := NewSynchronizedMap(l)

	var wg sync.WaitGroup

	// populate the map
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				assert.Nil(t, m.Put(k, v))
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()

	assert.Equal(t, n*n, m.Size(), "All items should be added")

	// remove the mappings
	base = 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				m.Remove(k)
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()

	assert.Equal(t, 0, m.Size(), "All items should be removed")
}

func TestSynchronizedMapContains(t *testing.T) {
	const n = 100
	const l = n * n << 2

	m := NewSynchronizedMap(l)

	var wg sync.WaitGroup

	// populate the map
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				assert.Nil(t, m.Put(k, v))
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()

	assert.Equal(t, n*n, m.Size(), "All items should be added")

	// check the mappings
	base = 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				assert.True(t, m.Contains(k))
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()
}

func TestSynchronizedMapKeys(t *testing.T) {
	const n = 100
	const l = n * n << 2

	m := NewSynchronizedMap(l)

	// populate the map
	var wg sync.WaitGroup
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := l + i + b
				assert.Nil(t, m.Put(k, v))
			}
			wg.Done()
		}(base)
		base += n
	}
	wg.Wait()

	assert.Equal(t, n*n, m.Size())

	keys := m.Keys()
	assert.Equal(t, n*n, len(keys))
}
