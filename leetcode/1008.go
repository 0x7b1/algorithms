package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(arr []int, i *int, bound int) *TreeNode {
	if len(arr) == *i || arr[*i] > bound {
		return nil
	}

	root := &TreeNode{Val: arr[*i]}
	*i++
	root.Left = build(arr, i, root.Val)
	root.Right = build(arr, i, bound)

	return root
}

func bstFromPreorder(preorder []int) *TreeNode {
	i := 0
	return build(preorder, &i, maxInt)
}

func printPreorder(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Val)
		printPreorder(t.Left)
		printPreorder(t.Right)
	}
}

func main() {
	t := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	printPreorder(t)
}
