package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightView(curr *TreeNode, result *[]int, currDepth int) {
	if curr == nil {
		return
	}

	if currDepth == len(*result) {
		*result = append(*result, curr.Val)
	}

	rightView(curr.Right, result, currDepth+1)
	rightView(curr.Left, result, currDepth+1)
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	rightView(root, &result, 0)
	return result
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{5, nil, nil}},
		&TreeNode{3, nil, &TreeNode{4, nil, nil}}}

	fmt.Println(rightSideView(root))
}
