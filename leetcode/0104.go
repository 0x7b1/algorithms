package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	depth := math.Max(float64(left), float64(right))

	return 1 + int(depth)
}

func main() {
	fmt.Println(maxDepth(
		&TreeNode{
			3,
			&TreeNode{
				9, nil, nil,
			},
			&TreeNode{
				20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil},
			},
		}))
}
