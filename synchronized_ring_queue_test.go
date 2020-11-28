package concurrent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSynchronizedRingQueue(t *testing.T) {
	for e := 2; e < 1024; e <<= 1 {
		t.Run("should not panic", func(t *testing.T) {
			_ = NewSynchronizedRingQueue(e)
		})
	}

	e := 2
	for n := 1; n < 1024; n++ {
		if n == e {
			e <<= 1
			continue
		}
		t.Run("should panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("did not panic: %d", n)
				}
			}()
			_ = NewSynchronizedRingQueue(n)
		})
	}
}

func TestSynchronizedRingQueueResize(t *testing.T) {
	for c := 2; c < 1024; c <<= 1 {
		t.Run("should be resized", func(t *testing.T) {
			q := NewSynchronizedRingQueue(c).(*SynchronizedRingQueue)
			assert.Equal(t, c, q.Capacity())
			for i := 0; i <= c; i++ {
				q.Offer(0)
			}
			assert.Equal(t, c<<1, q.Capacity())
		})
	}
}

func TestSynchronizedRingQueueOfferAndPoll(t *testing.T) {
	for c := 2; c < 1024; c <<= 1 {
		t.Run("should offer and poll", func(t *testing.T) {
			q := NewSynchronizedRingQueue(c)
			assert.Nil(t, q.Poll())
			for i := 0; i <= c; i++ {
				q.Offer(i)
			}
			for i := 0; i <= c; i++ {
				assert.Equal(t, i, q.Peek())
				assert.Equal(t, i, q.Poll())
			}
			assert.Nil(t, q.Peek())
			assert.Nil(t, q.Poll())
		})
	}
}

func TestSynchronizedRingQueueRange(t *testing.T) {
	const max = 1000

	for i := 1; i < max; i++ {
		q := NewSynchronizedRingQueue(2)
		for j := 0; j < i; j++ {
			q.Offer(j)
		}

		for k := 0; k < i/2; k++ {
			q.Poll()
		}

		var b []int
		q.Range(func(e interface{}) bool {
			b = append(b, e.(int))
			return true
		})
		assert.Equal(t, i-i/2, len(b))

		for j := 0; j < i-i/2; j++ {
			assert.Equal(t, j+i/2, b[j])
		}
	}
}
