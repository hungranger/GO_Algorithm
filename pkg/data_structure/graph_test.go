package data_structure

import (
	"testing"
)

func TestGraph_HasPathDFS(t *testing.T) {
	g := NewGraph()
	g = addNodes(g)

	type args struct {
		source      Node
		destination Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"Has path from A -> F",
			args{
				source:      Node{1, "A", nil},
				destination: Node{6, "F", nil},
			},
			true,
		},
		{
			"DO NOT has path from A -> G",
			args{
				source:      Node{1, "A", nil},
				destination: Node{7, "G", nil},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.HasPathDFS(tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("Graph.HasPathDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_HasPathBFS(t *testing.T) {
	g := NewGraph()
	g = addNodes(g)

	type args struct {
		source      Node
		destination Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"Has path from A -> F",
			args{
				source:      Node{1, "A", nil},
				destination: Node{6, "F", nil},
			},
			true,
		},
		{
			"DO NOT has path from A -> G",
			args{
				source:      Node{1, "A", nil},
				destination: Node{7, "G", nil},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.HasPathBFS(tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("Graph.HasPathBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func addNodes(g *Graph) *Graph {
	nA := Node{1, "A", nil}
	nB := Node{2, "B", nil}
	nC := Node{3, "C", nil}
	nD := Node{4, "D", nil}
	nE := Node{5, "E", nil}
	nF := Node{6, "F", nil}
	nG := Node{7, "G", nil}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddNode(&nG)

	g.AddEdge(nA, nB)
	g.AddEdge(nA, nC)
	g.AddEdge(nB, nE)
	g.AddEdge(nC, nE)
	g.AddEdge(nE, nF)
	g.AddEdge(nD, nA)
	g.AddEdge(nD, nF)

	return g
}
