package main

import "fmt"

type Node struct {
	visited   bool
	value     string
	neighbors []*Node
}

func (n Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func NewNode(value string) *Node {
	return &Node{
		false,
		value,
		nil,
	}
}

func (n *Node) connect(nodes ...*Node) {
	n.neighbors = append(n.neighbors, nodes...)
}

type Graph struct{}

func (g *Graph) DFS(n *Node) []*Node {
	stack := []*Node{n}
	visited := []*Node{}

	for len(stack) > 0 {
		e := stack[0]
		stack = stack[1:]
		e.visited = true

		visited = append([]*Node{e}, visited...)

		for _, v:= range e.neighbors {
			if !v.visited {
				stack = append([]*Node{v}, stack...)
			}
		}
	}

	return visited
}

func main() {
	v1 := NewNode("1")
	v2 := NewNode("2")
	v3 := NewNode("3")
	v4 := NewNode("4")
	v5 := NewNode("5")
	v6 := NewNode("6")
	v2.connect(v4, v5)
	v1.connect(v2, v3)
	v6.connect(v5)

	g := Graph{}
	fmt.Println(g.DFS(v1))
	fmt.Println(g.DFS(v6))
}
