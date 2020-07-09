package main

import "fmt"

type Node struct {
	value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.value)
}

type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}

	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}

	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (q *Queue) String() string {
	return fmt.Sprint(q.nodes)
}

func main() {
	Q := NewQueue(4)
	Q.Push(&Node{10})
	Q.Push(&Node{20})
	fmt.Println("->", Q.Pop())
	fmt.Println("->", Q.Pop())
	Q.Push(&Node{30})
	fmt.Println("->", Q.Pop())
	fmt.Println(Q)
}
