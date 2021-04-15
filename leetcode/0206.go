package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var behind *ListNode

	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
 		head = next
	}

	return behind
}

func (l ListNode) String() string {
	return fmt.Sprintf("%v -> %v", l.Val, l.Next)
}

func main() {
	list := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}}

	fmt.Println(reverseList(list))
}
