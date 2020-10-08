package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pushLeft(s *[]*TreeNode, n *TreeNode) {
	for n != nil {
		*s = append(*s, n)
		n = n.Left
	}
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := []int{}
	s1 := []*TreeNode{root1}
	s2 := []*TreeNode{root2}

	for !(len(s1) == 0) || !(len(s2) == 0) {
		s := []*TreeNode{}
		if len(s1) == 0 {
			s = s2
		} else if len(s2) == 0 {
			s = s1
		} else if s1[len(s1)-1].Val < s2[len(s2)-1].Val {
			s = s1
		} else {
			s = s2
		}

		n := s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, n.Val)
		pushLeft(&s, n)
	}

	return res
}

func ttt(s *[]int) {
	*s = append(*s, 10)
}

func main() {
	x := []int{1, 2, 3}
	pp := x
	ttt(&pp)
	fmt.Println(x)
	fmt.Println(pp)

	y := []int{1}
	z := y
	y = append(y, 2)
	fmt.Println(y, z)

	t1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{3, nil, nil}}

	fmt.Println(getAllElements(t1, t2))
}
