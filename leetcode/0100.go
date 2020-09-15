package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	} else if p == nil && q == nil {
		return true
	}
	return false
}

func main() {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(isSameTree(t1, t2))

	t3 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	t4 := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	fmt.Println(isSameTree(t3, t4))
}
