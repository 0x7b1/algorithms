package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			return 1
		}

		if root.Left == nil {
			return 1 + minDepth(root.Right)
		}

		if root.Right == nil {
			return 1 + minDepth(root.Left)
		}

		dLeft := minDepth(root.Left)
		dRight := minDepth(root.Right)

		if dLeft > dRight {
			return 1 + dRight
		} else {
			return 1 + dLeft
		}
	}

	return 0
}

func main() {
	t1 := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(minDepth(t1))

	t2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	fmt.Println(minDepth(t2))
}
