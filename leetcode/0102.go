package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := make([]int, 0)
		count := len(queue)

		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}

func main() {
	t :=
		&TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}

	fmt.Println(levelOrder(t))
}
