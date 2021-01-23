package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	var n1, n2, carry int

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		} else {
			n1 = 0
		}

		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		} else {
			n2 = 0
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return head.Next
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	fmt.Println(addTwoNumbers(l1, l2))
}
