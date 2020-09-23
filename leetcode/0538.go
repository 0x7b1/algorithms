package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	if root.Right != nil {
		inOrder(root.Right, sum)
	}

	*sum += root.Val
	root.Val = *sum

	if root.Left != nil {
		inOrder(root.Left, sum)
	}
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	inOrder(root, &sum)
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
	t := &TreeNode{5,
		&TreeNode{2, nil, nil},
		&TreeNode{13, nil, nil},
	}

	convertBST(t)
	printTree(t)
}
