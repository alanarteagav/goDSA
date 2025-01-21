package binaryTree

import ds "goDSA/pkg/dataStructures"

type BinaryTree[T ds.Ordered] interface {
	PreorderTraversal() []T
	InorderTraversal() []T
	PostorderTraversal() []T
}

type TreeNode[T ds.Ordered] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}
