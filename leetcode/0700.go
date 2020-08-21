package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	tmp := root

	for tmp != nil {
		if val > tmp.Val {
			tmp = tmp.Right
		} else if val < tmp.Val {
			tmp = tmp.Left
		} else if val == tmp.Val {
			break
		}
	}

	return tmp
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
		4, &TreeNode{
			2, &TreeNode{
				1, nil, nil,
			}, &TreeNode{
				3, nil, nil,
			},
		}, &TreeNode{
			7, nil, nil,
		},
	}

	t1 := searchBST(t, 2)
	fmt.Println(t1)
}