package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func helper(n *Node, arr []int) []int {
	if n != nil {
		for _, c := range n.Children {
			arr = helper(c, arr)
		}

		return append(arr, n.Val)
	} else {
		return arr
	}
}

func postorder(root *Node) []int {
	return helper(root, []int{})
}

func main() {
	t := &Node{1, []*Node{
		{3, []*Node{
			{5, nil},
			{6, nil},
		}},
		{2, nil},
		{4, nil},
	}}

	fmt.Println(postorder(t))
}
