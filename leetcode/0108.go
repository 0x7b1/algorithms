package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return &TreeNode{
		nums[len(nums)/2],
		sortedArrayToBST(nums[:len(nums)/2]),
		sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

func preOrderPrint(root *TreeNode) {
	if root != nil {
		preOrderPrint(root.Left)
		fmt.Println(root.Val)
		preOrderPrint(root.Right)
	}
}

func main() {
	bst := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	preOrderPrint(bst)
}
