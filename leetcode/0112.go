package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	t := &TreeNode{5,
		&TreeNode{4,
			&TreeNode{11,
				&TreeNode{7, nil, nil},
				&TreeNode{2, nil, nil}},
			nil},
		&TreeNode{8,
			&TreeNode{13, nil, nil},
			&TreeNode{4,
				nil,
				&TreeNode{1, nil, nil}}},
	}

	fmt.Println(hasPathSum(t, 22))
}
