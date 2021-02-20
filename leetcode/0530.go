package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min := int(1e10)
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if pre != nil {
				if cur.Val-pre.Val < min {
					min = cur.Val - pre.Val
				}
			}

			pre = cur
			cur = cur.Right
		}
	}

	return min
}

func main() {
	t := &TreeNode{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, nil}}
	fmt.Println(getMinimumDifference(t))
}
