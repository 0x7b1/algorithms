package main

import (
	"errors"
	"fmt"
	"log"
)

// Basic Structs

type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	return &Queue{[]int{}}
}

func (q *Queue) Push(n int) {
	q.arr = append(q.arr, n)
}

func (q *Queue) Pop() int {
	if len(q.arr) == 0 {
		return -1
	}

	n := q.arr[0]
	q.arr = q.arr[1:]

	return n
}

func (q *Queue) isEmpty() bool {
	return len(q.arr) == 0
}

type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack{[]int{}}
}

func (s *Stack) Push(n int)  {
	s.arr = append(s.arr, n)
}

func (s *Stack) Pop() int {
	if len(s.arr) == 0 {
		return -1
	}

	n := s.arr[len(s.arr) - 1]
	s.arr = s.arr[:len(s.arr) - 1]

	return n
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

// Graph Structs

type Node struct {
	id      int
	visited bool
	outs    []*Node
}

func NewNode(id int) *Node {
	return &Node{id, false, []*Node{}}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.id)
}

type Graph struct {
	nodes []*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: []*Node{},
	}
}

func (g *Graph) AddNode(id int) {
	n := NewNode(id)
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(a, b int) {
	n, err := g.findNode(a)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range n.outs {
		if v.id == b {
			log.Fatal("The connection already exists", a, b)
		}
	}

	m, err := g.findNode(b)
	if err != nil {
		log.Fatal(err)
	}

	n.outs = append(n.outs, m)
}

func (g *Graph) findNode(id int) (*Node, error) {
	for _, n := range g.nodes {
		if n.id == id {
			return n, nil
		}
	}

	return nil, errors.New("Node not found")
}

func (g *Graph) CleanState() {
	for _, n := range g.nodes {
		n.visited = false
	}
}

func (g *Graph) DFS(n *Node) []*Node {
	g.CleanState()
	q := NewStack()
	q.Push(n.id)

	for !q.isEmpty() {
		e, _ := g.findNode(q.Pop())
		e.visited = true
		fmt.Println("->", e.id)

		for _, v := range e.outs {
			if !v.visited {
				q.Push(v.id)
			}
		}
	}

	return nil
}

func (g *Graph) String() string {
	out := ""

	for _, n := range g.nodes {
		out += fmt.Sprintf("%v => %v\n", n, n.outs)
	}

	return out
}

func main() {
	g := NewGraph()
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	//g.AddEdge(5, 3)
	g.AddEdge(5, 6)

	n1, _ := g.findNode(1)
	g.DFS(n1)

	fmt.Println(g)
}
