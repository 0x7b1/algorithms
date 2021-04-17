package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var node *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = node
		node = slow
		slow = next
	}

	for node != nil {
		if node.Val != head.Val {
			return false
		}

		node = node.Next
		head = head.Next
	}

	return true
}

func (l ListNode) String() string {
	return fmt.Sprintf("%v -> %v", l.Val, l.Next)
}

func main() {
	list := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{2,
					&ListNode{1, nil}}}}}

	fmt.Println(isPalindrome(list))

	list1 := &ListNode{1,
		&ListNode{2,
				&ListNode{2,
					&ListNode{1, nil}}}}

	fmt.Println(isPalindrome(list1))

	list2 := &ListNode{1,
		&ListNode{2, nil}}

	fmt.Println(isPalindrome(list2))
}
