package data_structure

import (
	"fmt"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	// type args []int
	// tests := []struct {
	// 	name string
	// 	bst  *BST
	// 	args args
	// 	want
	// }{
	// 	// TODO: Add test cases.
	// 	{
	// 		"Insert to a nil BST",
	// 		NewBST(),
	// 		args{6,7,8,4,7,9,2,1,3},

	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		tt.bst.Insert(tt.args.value)
	// 	})
	// }

	bst := NewBST()
	bst.Insert(7)
	bst.Insert(20)
	bst.Insert(9)
	bst.Insert(8)
	bst.Insert(1)

	fmt.Printf("BTS 2: %v\n", bst.Contains(2))
	fmt.Printf("BST 9: %v\n", bst.Contains(9))
	fmt.Printf("DFS 2: %v\n", bst.ContainsDFS(2))
	fmt.Printf("DFS 9: %v\n", bst.ContainsDFS(9))
	fmt.Printf("BFS 2: %v\n", bst.ContainsBFS(2))
	fmt.Printf("BFS 9: %v\n", bst.ContainsBFS(9))

	bst.PrintInOrder()
	isValidBST(bst.root)

	t.Error("List.Insert() error")

}
