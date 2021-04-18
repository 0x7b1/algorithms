package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(node *TreeNode, c chan int) {
	if node.Left == nil && node.Right == nil {
		c <- node.Val
		return
	}

	if node.Left != nil {
		travel(node.Left, c)
	}

	if node.Right != nil {
		travel(node.Right, c)
	}
}

func leafChan(root *TreeNode) chan int {
	c := make(chan int)

	go func() {
		travel(root, c)
		close(c)
	}()

	return c
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ch1 := leafChan(root1)
	ch2 := leafChan(root2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 && ok2 {
			if v1 == v2 {
				continue
			} else {
				return false
			}
		} else if !ok1 && !ok2 {
			return true
		} else {
			return false
		}
	}
}

func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(printTree(root1), printTree(root2))
}

func printTree(t *TreeNode) []int {
	stack := []*TreeNode{t}
	res := []int{}

	for len(stack) > 0 {
		indexLast := len(stack) - 1
		node := stack[indexLast]
		stack = stack[:indexLast]

		if node != nil {
			if node.Left == nil && node.Right == nil {
				res = append(res, node.Val)
			}

			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}

	return res
}

func main() {
	tree1 := &TreeNode{3,
		&TreeNode{5,
			&TreeNode{6, nil, nil},
			&TreeNode{2,
				&TreeNode{7, nil, nil},
				&TreeNode{4, nil, nil},
			},
		},
		&TreeNode{1,
			&TreeNode{9, nil, nil},
			&TreeNode{8, nil, nil},
		},
	}

	tree2 := &TreeNode{3,
		&TreeNode{5,
			&TreeNode{6, nil, nil},
			&TreeNode{7, nil, nil}},
		&TreeNode{1,
			&TreeNode{4, nil, nil},
			&TreeNode{2,
				&TreeNode{9, nil, nil},
				&TreeNode{8, nil, nil}}},
	}

	fmt.Println(leafSimilar2(tree1, tree2))
	fmt.Println(leafSimilar(tree1, tree2))
}
