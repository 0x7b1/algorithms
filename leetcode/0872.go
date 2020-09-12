package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(printTree(root1), printTree(root2))
}

func printTree(t *TreeNode) []int {
	stack := []*TreeNode{t}
	res := []int{}

	for len(stack) > 0 {
		indexLast := len(stack) - 1
		node := stack[indexLast]
		stack = stack[:indexLast]

		if node != nil {
			if node.Left == nil && node.Right == nil {
				res = append(res, node.Val)
			}

			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}

	return res
}

func printTree2(t *TreeNode) {
	if t != nil {
		printTree(t.Left)
		if t.Left == nil && t.Right == nil {
			fmt.Println(t.Val)
		}
		printTree(t.Right)
	}
}

func main() {
	tree1 := &TreeNode{
		3,
		&TreeNode{
			5,
			&TreeNode{6, nil, nil},
			&TreeNode{
				2,
				&TreeNode{7, nil, nil},
				&TreeNode{4, nil, nil},
			},
		},
		&TreeNode{
			1,
			&TreeNode{9, nil, nil},
			&TreeNode{8, nil, nil},
		},
	}

	//tree2 := &TreeNode{}
	fmt.Println(leafSimilar(tree1, tree1))
}
