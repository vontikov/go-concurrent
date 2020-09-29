package concurrent

import (
	"testing"

	"sort"
	"sync"

	"github.com/stretchr/testify/assert"
)

func TestSynchronizedSetEmpty(t *testing.T) {
	assert.Equal(t, 0, NewSynchronizedSet(0).Size(), "Empty set size must be equal to 0")
	assert.Equal(t, 0, NewSynchronizedSet(1000).Size(), "Empty set size must be equal to 0")
}

func TestSynchronizedSetAddContains(t *testing.T) {
	const n = 100

	set := NewSynchronizedSet(n / 2)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < n; i++ {
				set.Add(i)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, n, set.Size(), "Only unique items should be added")

	for i := 0; i < n; i++ {
		assert.True(t, set.Contains(i), "Should contain all unique items")
	}
}

func TestSynchronizedSetClear(t *testing.T) {
	const n = 100

	set := NewSynchronizedSet(n)
	assert.Equal(t, 0, set.Size(), "Should be empty")

	for i := 0; i < n; i++ {
		set.Add(i)
	}
	assert.Equal(t, n, set.Size(), "Should not be empty")

	set.Clear()
	assert.Equal(t, 0, set.Size(), "Should be empty")
}

func TestSynchronizedSetAddRange(t *testing.T) {
	const n = 100

	s := NewSynchronizedSet(n / 2)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < n; i++ {
				s.Add(i)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, n, s.Size())

	var elms []int
	s.Range(func(e interface{}) bool {
		elms = append(elms, e.(int))
		return true
	})

	sort.Slice(elms, func(i, j int) bool { return elms[i] < elms[j] })

	for i := 0; i < n; i++ {
		assert.Equal(t, i, elms[i], "Keys should be unique")
	}
}

func TestSynchronizedSetRemove(t *testing.T) {
	const n = 100

	s := NewSynchronizedSet(n / 2)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < n; i++ {
				s.Add(i)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, n, s.Size())

	for i := 0; i < n; i++ {
		s.Remove(i)
	}

	assert.Equal(t, 0, s.Size())
}
