package concurrent

import (
	"sync"
)

// SynchronizedRingQueue is a safe for concurrent use Queue implementation
type SynchronizedRingQueue struct {
	sync.RWMutex
	buf               []interface{}
	head, tail, count int
}

// NewSynchronizedRingQueue returns pointer to a new SynchronizedRingQueue instance
func NewSynchronizedRingQueue(initialCapacity int) Queue {
	if (initialCapacity < 2) || ((initialCapacity & (initialCapacity - 1)) != 0) {
		panic("initial capacity must be power of 2")
	}
	return &SynchronizedRingQueue{
		buf: make([]interface{}, initialCapacity),
	}
}

// Size implements Queue.Size
func (q *SynchronizedRingQueue) Size() int {
	q.RLock()
	defer q.RUnlock()
	return q.count
}

// Clear implements Queue.Clear
func (q *SynchronizedRingQueue) Clear() {
	q.Lock()
	defer q.Unlock()
	q.head = 0
	q.tail = 0
	q.count = 0
}

// Offer implements Queue.Offer
func (q *SynchronizedRingQueue) Offer(e interface{}) {
	q.Lock()
	defer q.Unlock()

	if q.count == len(q.buf) {
		q.resize()
	}

	q.buf[q.tail] = e
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

// Poll implements Queue.Poll
func (q *SynchronizedRingQueue) Poll() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.count <= 0 {
		return nil
	}

	ret := q.buf[q.head]
	q.buf[q.head] = nil
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--
	return ret
}

// Peek implements Queue.Peek
func (q *SynchronizedRingQueue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.count <= 0 {
		return nil
	}

	return q.buf[q.head]
}

// Capacity implements Queue.Capacity
func (q *SynchronizedRingQueue) Capacity() int {
	q.RLock()
	defer q.RUnlock()
	return len(q.buf)
}

func (q *SynchronizedRingQueue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}
