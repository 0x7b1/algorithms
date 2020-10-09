package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func collect(root *TreeNode, values *[]int) {
	if root != nil {
		collect(root.Left, values)
		*values = append(*values, root.Val)
		collect(root.Right, values)
	}
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var values []int
	collect(root1, &values)
	collect(root2, &values)

	sort.Ints(values)

	return values
}

func main() {
	t1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{3, nil, nil}}

	fmt.Println(getAllElements(t1, t2))
}
