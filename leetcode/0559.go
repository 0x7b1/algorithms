package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	maxNode := float64(1)
	for i := 0; i < len(root.Children); i++ {
		maxNode = math.Max(float64(maxNode), 1 + float64(maxDepth(root.Children[i])))
	}

	return int(maxNode)
}

func main() {
	fmt.Println(maxDepth(&Node{1, []*Node{&Node{4, nil}}}))
}
