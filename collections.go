package concurrent

// Collection defines generic collection
type Collection interface {
	Size() int // Size returns the collection size
	Clear()    // Clear clears the collection
}

// List defines ordered collection of elements which may contain duplicates
type List interface {
	Collection

	Add(element interface{})   // Add adds a new element into the List
	Get(index int) interface{} // Get returns the element specified by its index
}

// Set defines unordered collection of element which may not contain duplicates
type Set interface {
	Collection

	Add(interface{})
	Contains(interface{}) bool
}

// Queue defines ordered in FIFO manner collection of elements which may contain
// duplicates
type Queue interface {
	Collection

	Offer(element interface{})
	Poll() interface{}
	Peek() interface{}
	Capacity() int
}

// Map represents a collection of key-value pairs
type Map interface {
	Collection

	// Put put a new key-value pair into the Map
	// if the key-value pair already exists it overwrites the existing value
	// with the new value
	Put(key interface{}, value interface{})

	// Get returns the value specified by the key if the key-value pair is
	// present, othervise returns nil
	Get(key interface{}) interface{}

	// PutIfAbsent puts the key-value pair (and returns true)
	// only if the key is absent, otherwise it returns false
	PutIfAbsent(key interface{}, value interface{}) bool
}
