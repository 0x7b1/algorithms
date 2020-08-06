package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode2(head *ListNode) *ListNode {
	heady := head

	length := 1
	for heady.Next != nil {
		heady = heady.Next
		length++
	}

	length /= 2

	for length > 0 {
		head = head.Next
		length--
	}

	return head
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}

func main() {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	mn := middleNode(head)

	fmt.Println(mn)
}
