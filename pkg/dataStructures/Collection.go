package dataStructures

type Collection[T Ordered] interface {
	// Adds an item to the collection
	Add(item T)
	// Removes an item from the collection
	Remove() T
	// Checks if the collection is empty
	IsEmpty() bool
}
