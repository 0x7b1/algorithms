package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}

	return a
}

func main() {
	c := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	a := &ListNode{4, &ListNode{1, c}}
	b := &ListNode{5, &ListNode{6, &ListNode{1, c}}}

	fmt.Println(getIntersectionNode(a, b))
}
