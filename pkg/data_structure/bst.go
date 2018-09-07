package data_structure

//Binary Search tree
type BST struct {
	root node
}

type node struct {
	value       interface{}
	left, right *node
}
