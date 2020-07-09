package main

import "fmt"

type Node struct {
	value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.value)
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}

	s.count--
	return s.nodes[s.count]
}

func main() {
	S := NewStack()
	S.Push(&Node{10})
	S.Push(&Node{20})
	fmt.Println(S.Pop())
	fmt.Println(S.Pop())
}
