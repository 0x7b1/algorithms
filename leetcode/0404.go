package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left != nil && current.Left.Left == nil && current.Left.Right == nil {
			sum += current.Left.Val
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return sum
}

func helper(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	return helper(root.Left, true) + helper(root.Right, false)
}

func sumOfLeftLeavesRecursive(root *TreeNode) int {
	return helper(root, false)
}

func main() {
	t := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}

	fmt.Println(sumOfLeftLeaves(t))
	fmt.Println(sumOfLeftLeavesRecursive(t))
}
