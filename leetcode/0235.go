package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

func main() {
	p := &TreeNode{2,
		&TreeNode{0, nil, nil},
		&TreeNode{4,
			&TreeNode{3, nil, nil},
			&TreeNode{5, nil, nil}}}
	q :=&TreeNode{8,
			&TreeNode{7, nil, nil},
			&TreeNode{9, nil, nil}}
	root := &TreeNode{6,
		p,
		q}

	pSearch := &TreeNode{2, nil, nil}
	qSearch := &TreeNode{4, nil, nil}

	fmt.Println(lowestCommonAncestor(root, pSearch, qSearch))
}
