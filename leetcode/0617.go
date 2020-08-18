package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

func printTree(t *TreeNode) {
	if t != nil {
		printTree(t.Left)
		printTree(t.Right)
		fmt.Println(t.Val)
	}
}

func main() {
	t1 := &TreeNode{
		1,
		&TreeNode{
			3,
			&TreeNode{
				5,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}

	t2 := &TreeNode{
		2,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	t12 := mergeTrees(t1, t2)
	printTree(t12)
}
