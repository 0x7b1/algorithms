package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative approach
func sumRootToLeaf1(root *TreeNode) int {
	sum := 0
	type nodeSum struct {
		node    *TreeNode
		currSum int
	}
	stack := []nodeSum{{root, 0}}

	for len(stack) > 0 {
		nodeTmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := nodeTmp.node
		currSum := nodeTmp.currSum

		if node != nil {
			currSum = (currSum << 1) | node.Val

			if node.Left == nil && node.Right == nil {
				sum += currSum
			} else {
				stack = append(stack, nodeSum{node.Left, currSum})
				stack = append(stack, nodeSum{node.Right, currSum})
			}
		}
	}

	return sum
}

func preorder(node *TreeNode, currSum int, totalSum *int) {
	if node != nil {
		currSum = (currSum << 1) | node.Val

		if node.Left == nil && node.Right == nil {
			*totalSum += currSum
		} else {
			preorder(node.Left, currSum, totalSum)
			preorder(node.Right, currSum, totalSum)
		}
	}
}

// Recursive Approach
func sumRootToLeaf2(root *TreeNode) int {
	sum := 0
	preorder(root, 0, &sum)

	return sum
}

// Morris preorder traversal
func sumRootToLeaf(root *TreeNode) int {
	rootToLeaf, currNumber := 0, 0

	for root != nil {
		// If there is a left child, then compute the predecessor
		// If there's no link predecessor.right = root then set it
		// If there's a link predecessor.right = root then break it

		if root.Left != nil {
			// Predecessor node is one step to the left
			// and then to the right
			predecessor := root.Left
			steps := 1
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
				steps++
			}

			// Set link predecessor.right = root
			// and go to explore the left subtree
			if predecessor.Right == nil {
				currNumber = (currNumber << 1) | root.Val
				predecessor.Right = root
				root = root.Left
			} else {
				// Break the link predecessor.right = root
				// Once the link is broken,
				// it's time to change subtree and go to the right

				// If you're on the leaf, update the sum
				if predecessor.Left == nil {
					rootToLeaf += currNumber
				}

				// This part of the tree is explored, backtrack
				for i := 0; i < steps; i++ {
					currNumber >>= 1
				}

				predecessor.Right = nil
				root = root.Right
			}
		} else {
			// If there's no left child
			// Then just go right
			currNumber = (currNumber << 1) | root.Val
			// If you're on the leaf, update the sum
			if root.Right == nil {
				rootToLeaf += currNumber
			}

			root = root.Right
		}
	}

	return rootToLeaf
}

func main() {
	tree := &TreeNode{
		1,
		&TreeNode{
			0,
			&TreeNode{
				0,
				nil,
				nil,
			},
			&TreeNode{
				1,
				nil,
				nil,
			},
		},
		&TreeNode{
			1,
			&TreeNode{
				0,
				nil,
				nil,
			},
			&TreeNode{
				1,
				nil,
				nil,
			},
		},
	}

	fmt.Println(sumRootToLeaf(tree))
}
