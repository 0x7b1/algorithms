package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := []*TreeNode{}

	for i := 0; i < len(nums); i++ {
		curr := &TreeNode{nums[i], nil, nil}

		for len(stack) > 0 && stack[len(stack) - 1].Val < nums[i] {
			curr.Left = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}

		if len(stack) > 0 {
			stack[len(stack) - 1].Right = curr
		}

		stack = append(stack, curr)
	}

	return stack[0]
}

func printTree(t *TreeNode) {
	if t != nil {
		printTree(t.Left)
		fmt.Println(t.Val)
		printTree(t.Right)
	}
}

func main() {
	printTree(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
