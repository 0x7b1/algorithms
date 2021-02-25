package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%v -> %v", l.Val, l.Next)
}

func main() {
	list := &ListNode{1,
		&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{3,
						nil}}}}}
	fmt.Println(deleteDuplicates(list))
}
