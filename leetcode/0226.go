package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%v\n(%v,%v)", t.Val, t.Left, t.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	r := root.Right
	l := root.Left

	root.Left = invertTree(r)
	root.Right = invertTree(l)

	return root
}

func main() {
	t := &TreeNode{
		4,
		&TreeNode{
			2,
			&TreeNode{
				1, nil, nil,
			}, &TreeNode{
				3, nil, nil,
			},
		}, &TreeNode{
			7,
			&TreeNode{
				6, nil, nil,
			},
			&TreeNode{
				9, nil, nil,
			},
		},
	}

	fmt.Println(t)

	ti := invertTree(t)
	fmt.Println(ti)
}
