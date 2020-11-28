package concurrent

// Equals indicates whether l is "equal to" r.
type Equals func(l, r interface{}) bool

// Collection defines generic collection.
type Collection interface {
	// Size returns the collection size
	Size() int

	// Clear clears the collection
	Clear()
}

// List defines ordered collection of elements which may contain duplicates.
type List interface {
	Collection

	// Add adds the element e into the List.
	Add(e interface{})

	// Get returns the element specified by its index i.
	Get(i int) interface{}

	// Remove the first occurrence of the element e from the list if it is present.
	Remove(e interface{}, eq Equals) bool

	// Range calls f sequentially for each element present in the list.
	// If f returns false, range stops the iteration.
	Range(f func(e interface{}) bool)
}

// Queue defines ordered in FIFO manner collection of elements which may contain
// duplicates.
type Queue interface {
	Collection

	// Offer inserts the element e into the Queue.
	Offer(e interface{})

	// Poll retrieves and removes the head of the Queue; returns nil if the
	// queue is empty.
	Poll() interface{}

	// Peek retrieves, but does not remove, the head of the queue; returns nil
	// if the queue is empty.
	Peek() interface{}

	// Range calls f sequentially for each element present in the queue.
	// If f returns false, range stops the iteration.
	Range(f func(e interface{}) bool)
}

// Set defines unordered collection of element which may not contain duplicates.
type Set interface {
	Collection

	// Add adds the element e into the set.
	Add(e interface{})

	// Contains returns true if the set contains the element e.
	Contains(e interface{}) bool

	// Range calls f sequentially for each element present in the set.
	// If f returns false, range stops the iteration.
	Range(f func(e interface{}) bool)

	// Remove removes the element e from the set if it is present.
	Remove(e interface{})
}

// Map represents a collection of key-value pairs.
type Map interface {
	Collection

	// Put puts a new key-value pair into the Map.
	// If the key already exists overwrites the existing value with the new one.
	// Returns the previous value associated with key, or nil if there was no mapping
	// for key.
	Put(k interface{}, v interface{}) interface{}

	// PutIfAbsent puts the key-value pair (and returns true)
	// only if the key is absent, otherwise it returns false.
	PutIfAbsent(k interface{}, v interface{}) bool

	// ComputeIfAbsent computes the mapping function f and inserts its value
	// (unless nil) under the key k, if the key does not exist and.
	// Returns the mapping and the flag indicating if the mapping was created.
	ComputeIfAbsent(k interface{}, f func() interface{}) (interface{}, bool)

	// Contains returns true if the map contains the key k.
	Contains(k interface{}) bool

	// Get returns the value specified by the key if the key-value pair is
	// present, othervise returns nil.
	Get(k interface{}) interface{}

	// Range calls f sequentially for each key and value present in the map.
	// If f returns false, range stops the iteration.
	Range(f func(k, v interface{}) bool)

	// Remove removes the key-value pair specified by the key k from the map
	// if it is present.
	Remove(k interface{})

	// Keys returns the keys contained in the map.
	Keys() []interface{}
}
