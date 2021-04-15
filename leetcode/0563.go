package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func postOrder(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	left := postOrder(root.Left, sum)
	right := postOrder(root.Right, sum)
	tilt := abs(left - right)

	*sum += tilt

	return left + right + root.Val
}

func findTilt(root *TreeNode) int {
	sum := 0
	postOrder(root, &sum)
	return sum
}

func main() {
	t := &TreeNode{6,
		&TreeNode{2,
			&TreeNode{0, nil, nil},
			&TreeNode{0, nil, nil}},
		&TreeNode{7,
			nil,
			&TreeNode{0, nil, nil}}}

	fmt.Println(findTilt(t))
}
