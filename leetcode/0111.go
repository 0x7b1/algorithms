package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthDFS(root *TreeNode) int {
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			currNode := queue[0]
			queue = queue[1:]
			if currNode.Left == nil && currNode.Right == nil {
				return level
			}

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}

			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}

		level++
	}

	return level
}

func main() {
	t1 := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(minDepth(t1))

	t2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	fmt.Println(minDepth(t2))
}
