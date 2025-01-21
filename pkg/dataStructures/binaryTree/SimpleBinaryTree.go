package binaryTree

import ds "github.com/alanarteagav/goDSA/pkg/dataStructures"

type Tree[T ds.Ordered] struct {
	root *TreeNode[T]
}

func NewTree[T ds.Ordered]() *Tree[T] {
	return &Tree[T]{
		root: nil,
	}
}

func (t *Tree[T]) PreorderTraversal() []T {
	traversal := []T{}
	var preorderDFS func(node *TreeNode[T])
	preorderDFS = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		traversal = append(traversal, node.value)
		preorderDFS(node.left)
		preorderDFS(node.right)
	}
	preorderDFS(t.root)
	return traversal
}

func (t *Tree[T]) InorderTraversal() []T {
	traversal := []T{}
	var inorderDFS func(node *TreeNode[T])
	inorderDFS = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		inorderDFS(node.left)
		traversal = append(traversal, node.value)
		inorderDFS(node.right)
	}
	inorderDFS(t.root)
	return traversal
}

func (t *Tree[T]) PostorderTraversal() []T {
	traversal := []T{}
	var postorderDFS func(node *TreeNode[T])
	postorderDFS = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		postorderDFS(node.left)
		postorderDFS(node.right)
		traversal = append(traversal, node.value)
	}
	postorderDFS(t.root)
	return traversal
}
