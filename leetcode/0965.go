package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		node := stack[lastIndex]
		stack = stack[:lastIndex]

		if node != nil {
			if node.Val != val {
				return false
			}
			stack = append(stack, node.Left, node.Right)
		}
	}

	return true
}

func main() {
	tree := &TreeNode{
		1,
		&TreeNode{
			1,
			&TreeNode{
				1, nil, nil,
			},
			&TreeNode{
				1, nil, nil,
			},
		},
		&TreeNode{
			1,
			nil,
			&TreeNode{1, nil, nil},
		},
	}

	fmt.Println(isUnivalTree(tree))
}
