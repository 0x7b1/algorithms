package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, path string, paths *[]string) {
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, fmt.Sprintf("%v%v", path, root.Val))
	}

	if root.Left != nil {
		searchBST(root.Left, fmt.Sprintf("%v%v->", path, root.Val), paths)
	}

	if root.Right != nil {
		searchBST(root.Right, fmt.Sprintf("%v%v->", path, root.Val), paths)
	}
}

func binaryTreePathsRecursive(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var paths []string
	searchBST(root, "", &paths)
	return paths
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var paths []string
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	leftPaths := binaryTreePaths(root.Left)
	for _, path := range leftPaths {
		paths = append(paths, strconv.Itoa(root.Val) + "->" + path)
	}

	rightPaths := binaryTreePaths(root.Right)
	for _, path := range rightPaths {
		paths = append(paths, strconv.Itoa(root.Val) + "->" + path)
	}

	return paths
}

func main() {
	t := &TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil}}

	fmt.Println(binaryTreePaths(t))
}
