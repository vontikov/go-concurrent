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
	const big = n * n << 2

	m := NewSynchronizedMap(big)

	var wg sync.WaitGroup
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := big + i + b
				m.Put(k, v)
			}
			wg.Done()
		}(base)
		base += n
	}

	wg.Wait()

	assert.Equal(t, n*n, m.Size(), "All items should be added")

	// extract and sort
	s := make([]int, 0, m.Size())
	for i := 0; i < n*n; i++ {
		s = append(s, m.Get(i).(int))
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	for i := 0; i < n*n; i++ {
		assert.Equal(t, big+i, s[i], "Items should be unique")
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
	const big = n * n << 2

	m := NewSynchronizedMap(big)

	// first pass
	var wg sync.WaitGroup
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := big + i + b
				r := m.PutIfAbsent(k, v)
				assert.True(t, r, "Should add the pair")
			}
			wg.Done()
		}(base)
		base += n
	}

	wg.Wait()
	assert.Equal(t, n*n, m.Size(), "All items should be added")

	// extract and sort
	s := make([]int, 0, m.Size())
	for i := 0; i < n*n; i++ {
		s = append(s, m.Get(i).(int))
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	for i := 0; i < n*n; i++ {
		assert.Equal(t, big+i, s[i], "Items should be unique")
	}

	// second pass
	base = 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				k := i + b
				v := big + i + b
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