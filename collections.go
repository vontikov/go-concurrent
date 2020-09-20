package concurrent

// Collection defines generic collection
type Collection interface {
	// Size returns the collection size
	Size() int
	// Clear clears the collection
	Clear()
}

// List defines ordered collection of elements which may contain duplicates
type List interface {
	Collection
	// Add adds the element e into the List
	Add(e interface{})
	// Get returns the element specified by its index i
	Get(i int) interface{}
}

// Set defines unordered collection of element which may not contain duplicates
type Set interface {
	Collection
	// Add adds the element e into the Set
	Add(e interface{})
	// Contains returns true if the Set contains the specified element e
	Contains(e interface{}) bool
}

// Queue defines ordered in FIFO manner collection of elements which may contain
// duplicates
type Queue interface {
	Collection
	// Offer inserts the element e into the Queue
	Offer(e interface{})
	// Poll retrieves and removes the head of the Queue; returns nil if the Queue is empty
	Poll() interface{}
	// Peek retrieves, but does not remove, the head of the queue; returns nil if the Queue is empty
	Peek() interface{}
}

// Map represents a collection of key-value pairs
type Map interface {
	Collection

	// Put puts a new key-value pair into the Map.
	// If the key already exists overwrites the existing value with the new one.
	// Returns the previous value associated with key, or nil if there was no mapping
	// for key.
	Put(key interface{}, value interface{}) interface{}

	// Get returns the value specified by the key if the key-value pair is
	// present, othervise returns nil
	Get(key interface{}) interface{}

	// PutIfAbsent puts the key-value pair (and returns true)
	// only if the key is absent, otherwise it returns false
	PutIfAbsent(key interface{}, value interface{}) bool

	// Range calls f sequentially for each key and value present in the map.
	// If f returns false, range stops the iteration.
	Range(f func(key, value interface{}) bool)
}
