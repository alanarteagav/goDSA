package dataStructures

type Collection[T Ordered] interface {
	Add(item T)
	Remove() T
	IsEmpty() bool
}
