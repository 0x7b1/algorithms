package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, pre, res *int) {
	if root == nil {
		return
	}

	helper(root.Left, pre, res)
	*res = min(*res, root.Val - *pre)
	*pre = root.Val

	helper(root.Right, pre, res)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}


func minDiffInBST(root *TreeNode) int {
	pre, res := -int(1e10), int(1e10)
	helper(root, &pre, &res)
	return res
}

func main() {
	t := &TreeNode{1,
		nil,
		&TreeNode{3,
			&TreeNode{2, nil, nil},
			nil}}

	fmt.Println(minDiffInBST(t))
}
