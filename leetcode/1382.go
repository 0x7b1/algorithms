package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraverse(sortedArr *[]*TreeNode, root *TreeNode) {
	if root != nil {
		inOrderTraverse(sortedArr, root.Left)
		*sortedArr = append(*sortedArr, root)
		inOrderTraverse(sortedArr, root.Right)
	}
}

func sortedArrayToBST(sortedArr []*TreeNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := sortedArr[mid]
	root.Left = sortedArrayToBST(sortedArr, start, mid-1)
	root.Right = sortedArrayToBST(sortedArr, mid+1, end)

	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	var sortedArr []*TreeNode
	inOrderTraverse(&sortedArr, root)
	balancedBST := sortedArrayToBST(sortedArr, 0, len(sortedArr)-1)

	return balancedBST
}

func (t *TreeNode) String() string {
	if t == nil {
		return "-"
	} else {
		return fmt.Sprintf("%v (%v,%v)", t.Val, t.Left, t.Right)
	}
}

func main() {
	t := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				&TreeNode{
					4,
					nil,
					nil,
				},
				nil,
			},
			nil,
		},
		nil,
	}

	tb := balanceBST(t)

	fmt.Println(tb)
}
