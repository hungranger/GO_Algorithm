package data_structure

import (
	"fmt"
	"sort"
)

//Binary Search tree
type BST struct {
	root *node
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Insert(value int) {
	if bst.root == nil {
		bst.root = newNode(value)
		return
	}

	bst.root.insert(value)
}

func (bst *BST) Contains(value int) bool {
	return bst.root.contains(value)
}

func (bst *BST) ContainsBFS(value int) bool {
	return bst.root.containsBFS(value)
}

func (bst *BST) ContainsDFS(value int) bool {
	return bst.root.containsDFS(value)
}

func (bst *BST) PrintInOrder() {
	bst.root.printInOrder()
}

type node struct {
	value       int
	left, right *node
}

func newNode(value int) *node {
	return &node{value, nil, nil}
}

func (n *node) insert(value int) {
	if value >= n.value {
		if n.right == nil {
			n.right = newNode(value)
		} else {
			n.right.insert(value)
		}
	} else {
		if n.left == nil {
			n.left = newNode(value)
		} else {
			n.left.insert(value)
		}
	}
}

// Depth First Search
func (n *node) contains(value int) bool {
	left := false
	right := false

	if n.value == value {
		return true
	}

	if n.value > value {
		if n.left != nil {
			left = n.left.contains(value)
		}
	} else {
		if n.right != nil {
			return n.right.contains(value)
		}
	}

	return left || right
}

func (n *node) containsBFS(value int) bool {
	queue := []*node{}
	queue = append(queue, n)
	for len(queue) > 0 {
		if queue[0].value == value {
			return true
		}

		if queue[0].left != nil {
			queue = append(queue, queue[0].left)
		}

		if queue[0].right != nil {
			queue = append(queue, queue[0].right)
		}

		queue[0] = nil
		queue = queue[1:]
	}
	return false
}

func (n *node) containsDFS(value int) bool {
	left := false
	right := false

	if n.value == value {
		return true
	}

	if n.left != nil {
		left = n.left.containsDFS(value)
	}

	if n.right != nil {
		right = n.right.containsDFS(value)
	}

	return left || right

}

func (n *node) printInOrder() {
	if n.left != nil {
		n.left.printInOrder()
	}
	fmt.Print(n.value)
	if n.right != nil {
		n.right.printInOrder()
	}
}

func isValidBST(root *node) bool {
	arr := []int{}
	arr = toArray(root)
	fmt.Println(arr)
	return sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func toArray(root *node) []int {
	l, m, r := []int{}, []int{}, []int{}

	if root.left != nil {
		l = toArray(root.left)
	}
	m = append(m, root.value)
	if root.right != nil {
		r = toArray(root.right)
	}

	l = append(l, m...)
	l = append(l, r...)

	return l
}
