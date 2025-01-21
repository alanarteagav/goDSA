package binaryTree

import ds "goDSA/pkg/dataStructures"

type BinarySearchTree[T ds.Ordered] struct {
	root *TreeNode[T]
}

func NewBinarySearchTree[T ds.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root: nil,
	}
}

func (t *BinarySearchTree[T]) AddIterative(value T) {
	newNode := new(TreeNode[T])
	newNode.value = value

	if t.root == nil {
		t.root = newNode
		return
	}
	var parent *TreeNode[T]
	node := t.root
	for node != nil {
		parent = node
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	if value < parent.value {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}

func (t *BinarySearchTree[T]) AddRecursive(value T) {
	t.root = addAux(t.root, value)
}

func addAux[T ds.Ordered](node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		n := new(TreeNode[T])
		n.value = value
		return n
	}
	if value < node.value {
		node.left = addAux(node.left, value)
	} else {
		node.right = addAux(node.right, value)
	}
	return node
}

func (t *BinarySearchTree[T]) PreorderTraversal() []T {
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

func (t *BinarySearchTree[T]) InorderTraversal() []T {
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

func (t *BinarySearchTree[T]) PostorderTraversal() []T {
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

func (t *BinarySearchTree[T]) Delete(value T) {
	t.root = deleteNode(t.root, value)
}

func deleteNode[T ds.Ordered](node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return node
	}
	if value < node.value {
		node.left = deleteNode(node.left, value)
	} else if value > node.value {
		node.right = deleteNode(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			getInorderSuccessor := func(node *TreeNode[T]) *TreeNode[T] {
				n := node.right
				for n != nil && n.left != nil {
					n = n.left
				}
				return n
			}
			successor := getInorderSuccessor(node)
			node.value = successor.value
			node.right = deleteNode(node.right, successor.value)
			return node
		}
	}
	return node
}
