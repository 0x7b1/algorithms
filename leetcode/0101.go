package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compareSymmetry(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val == right.Val {
			return compareSymmetry(left.Left, right.Right) &&
				compareSymmetry(left.Right, right.Left)
		}
	}

	if left == nil && right == nil {
		return true
	}

	return false
}

func isSymmetric(root *TreeNode) bool {
	if root != nil {
		return compareSymmetry(root.Left, root.Right)
	}

	return true
}

func main() {
	t1 := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{
			2,
			&TreeNode{4, nil, nil},
			&TreeNode{3, nil, nil},
		},
	}

	t2 := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{4, nil, nil},
		},
		&TreeNode{
			2,
			nil,
			&TreeNode{4, nil, nil},
		},
	}

	fmt.Println(isSymmetric(t1))
	fmt.Println(isSymmetric(t2))
}
