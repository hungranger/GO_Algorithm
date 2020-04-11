package data_structure

import (
	"testing"
)

func addBSTNodes(bst *BST) *BST {
	bst.Insert(7)
	bst.Insert(20)
	bst.Insert(9)
	bst.Insert(8)
	bst.Insert(1)
	bst.PrintInOrder()
	// isValidBST(bst.root)
	return bst
}

func TestBST_Contains(t *testing.T) {
	bst, tests := buildContainParams()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bst.Contains(tt.args); got != tt.want {
				t.Errorf("BST.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_ContainsBFS(t *testing.T) {
	bst, tests := buildContainParams()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bst.ContainsBFS(tt.args); got != tt.want {
				t.Errorf("BST.ContainsBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_ContainsDFS(t *testing.T) {
	bst, tests := buildContainParams()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bst.ContainsDFS(tt.args); got != tt.want {
				t.Errorf("BST.ContainsDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

type tests []struct {
	name string
	args int
	want bool
}

func buildContainParams() (bst *BST, t tests) {
	bst = NewBST()
	bst = addBSTNodes(bst)

	t = tests{
		// TODO: Add test cases.
		{
			"DO NOT contain value number 2",
			2,
			true,
		},
		{
			"DO contain value number 9",
			9,
			true,
		},
	}
	return
}
