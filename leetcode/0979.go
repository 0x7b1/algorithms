package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func dfsCoins(node *TreeNode, steps *int) int {
	if node == nil {
		return 0
	}

	l := dfsCoins(node.Left, steps)
	r := dfsCoins(node.Right, steps)

	*steps += abs(l) + abs(r)

	return node.Val + l + r - 1
}

func distributeCoins(root *TreeNode) int {
	steps := 0
	dfsCoins(root, &steps)
	return steps
}

func main() {
	t1 := &TreeNode{3,
		&TreeNode{0, nil, nil},
		&TreeNode{0, nil, nil},
	}

	t2 := &TreeNode{0,
		&TreeNode{3, nil, nil},
		&TreeNode{0, nil, nil},
	}

	t3 := &TreeNode{1,
		&TreeNode{0, nil, &TreeNode{3, nil, nil}},
		&TreeNode{0, nil, nil},
	}

	fmt.Println(distributeCoins(t1))
	fmt.Println(distributeCoins(t2))
	fmt.Println(distributeCoins(t3))
}
