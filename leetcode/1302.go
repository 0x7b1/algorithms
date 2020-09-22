package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helperDFS(root *TreeNode, level int, maxLevel, sum *int) {
	if root == nil {
		return
	}

	if level > *maxLevel {
		*maxLevel, *sum = level, root.Val
	} else if level == *maxLevel {
		*sum += root.Val
	}

	helperDFS(root.Left, level+1, maxLevel, sum)
	helperDFS(root.Right, level+1, maxLevel, sum)
}

func deepestLeavesSum(root *TreeNode) int {
	maxLevel, sum := 0, 0
	helperDFS(root, 0, &maxLevel, &sum)
	return sum
}

func main() {
	t := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				&TreeNode{7, nil, nil},
				nil,
			}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil,
			&TreeNode{6, nil, &TreeNode{8, nil, nil}}},
	}

	fmt.Println(deepestLeavesSum(t))
}
