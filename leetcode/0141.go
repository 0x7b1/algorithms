package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	var a, b ListNode

	a.Val = 3
	a.Next = &b
	b.Val = 2
	b.Next = &a

	fmt.Println(hasCycle(&a))
}
