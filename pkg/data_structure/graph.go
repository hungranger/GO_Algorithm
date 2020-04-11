package data_structure

import "fmt"

type Node struct {
	id       int
	value    interface{}
	adjacent []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *Node) AddAdjacent(node *Node) {
	n.adjacent = append(n.adjacent, node)
}

type Graph struct {
	nodeLookup map[int]*Node
}

func (g *Graph) String() {
	s := ""
	for _, node := range g.nodeLookup {
		s += node.String() + " -> "
		for _, child := range node.adjacent {
			s += child.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func NewGraph() *Graph {
	return &Graph{
		nodeLookup: map[int]*Node{},
	}
}

func (g *Graph) AddNode(node *Node) {
	g.nodeLookup[node.id] = node
}

func (g *Graph) getNode(id int) *Node {
	if node, ok := g.nodeLookup[id]; ok {
		return node
	}
	return nil
}

func (g *Graph) AddEdge(source Node, destination Node) {
	s := g.getNode(source.id)
	d := g.getNode(destination.id)
	s.AddAdjacent(d)
}

func (g *Graph) HasPathDFS(source Node, destination Node) bool {
	s := g.getNode(source.id)
	d := g.getNode(destination.id)
	visited := make(map[int]bool)
	return g.hasPathDFS(s, d, visited)
}

func (g *Graph) hasPathDFS(source *Node, destination *Node, visited map[int]bool) bool {
	if _, ok := visited[source.id]; ok {
		return false
	}
	visited[source.id] = true

	if source == destination {
		return true
	}

	for _, child := range source.adjacent {
		if g.hasPathDFS(child, destination, visited) {
			return true
		}
	}

	return false
}

func (g *Graph) HasPathBFS(source Node, destination Node) bool {
	s := g.getNode(source.id)
	d := g.getNode(destination.id)

	nextToVisit := []*Node{}
	visited := make(map[int]bool)
	nextToVisit = append(nextToVisit, s)
	for len(nextToVisit) > 0 {
		node := nextToVisit[0]
		if node == d {
			return true
		}

		if _, ok := visited[node.id]; ok {
			nextToVisit = nextToVisit[1:]
			continue
		}
		visited[node.id] = true

		for _, child := range node.adjacent {
			nextToVisit = append(nextToVisit, child)
		}

		nextToVisit = nextToVisit[1:]
	}
	return false
}
