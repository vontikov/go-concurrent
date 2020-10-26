package concurrent

import (
	"testing"

	"sort"
	"sync"

	"github.com/stretchr/testify/assert"
)

func TestSynchronizedListEmpty(t *testing.T) {
	assert.Equal(t, 0, NewSynchronizedList(0).Size(), "Empty list size must be equal to 0")
	assert.Equal(t, 0, NewSynchronizedList(1000).Size(), "Empty list size must be equal to 0")
}

func TestSynchronizedListAddGet(t *testing.T) {
	const n = 100

	list := NewSynchronizedList(n / 2)

	var wg sync.WaitGroup
	base := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(b int) {
			for i := 0; i < n; i++ {
				list.Add(i + b)
			}
			wg.Done()
		}(base)
		base += n
	}

	wg.Wait()

	assert.Equal(t, n*n, list.Size(), "All items should be added")

	// extract and sort
	s := make([]int, 0, list.Size())
	for i := 0; i < n*n; i++ {
		s = append(s, list.Get(i).(int))
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	for i := 0; i < n*n; i++ {
		assert.Equal(t, i, s[i], "Items should be unique")
	}
}

func TestSynchronizedListClear(t *testing.T) {
	const n = 100

	list := NewSynchronizedList(n)
	assert.Equal(t, 0, list.Size(), "Should be empty")

	for i := 0; i < n; i++ {
		list.Add(i)
	}
	assert.Equal(t, n, list.Size(), "Should not be empty")

	list.Clear()
	assert.Equal(t, 0, list.Size(), "Should be empty")
}

func TestSynchronizedListRemove(t *testing.T) {
	const n = 15

	list := NewSynchronizedList(n)
	for i := 0; i < n; i++ {
		list.Add(i)
	}
	assert.Equal(t, n, list.Size())

	eq := func(l, r interface{}) bool { return l.(int) == r.(int) }

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(e int) {
			assert.True(t, list.Remove(e, eq))
			assert.False(t, list.Remove(e, eq))
			wg.Done()
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 0, list.Size(), "Should be empty")
}

func TestSynchronizedListAddRange(t *testing.T) {
	const n = 100

	list := NewSynchronizedList(0)

	for i := 0; i < n; i++ {
		list.Add(i * i)
	}

	var elms []interface{}
	list.Range(func(e interface{}) bool {
		elms = append(elms, e)
		return true
	})
	for i := 0; i < n; i++ {
		assert.Equal(t, i*i, elms[i])
	}
}
