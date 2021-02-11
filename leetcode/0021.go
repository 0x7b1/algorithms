package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("%v - %v", n.Val, n.Next)
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	fmt.Println("BEFORE")
	fmt.Println("L1:", l1)
	fmt.Println("L2:", l2)

	l3 := mergeTwoLists(l1, l2)

	fmt.Println("AFTER")
	fmt.Println("L1:", l1)
	fmt.Println("L2:", l2)
	fmt.Println("L3:", l3)
}
