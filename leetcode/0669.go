package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	} else if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}

func printTree(t *TreeNode) {
	if t != nil {
		printTree(t.Left)
		fmt.Println(t.Val)
		printTree(t.Right)
	}
}

func main() {
	t1 := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}
	printTree(trimBST(t1, 1, 2))

	t2 := &TreeNode{3, &TreeNode{0, nil, &TreeNode{2, &TreeNode{1, nil, nil}, nil}}, &TreeNode{4, nil, nil}}
	printTree(trimBST(t2, 1, 3))
}
