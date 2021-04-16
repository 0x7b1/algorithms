package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	rev := &ListNode{}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		rev, rev.Next, slow = slow, rev, slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for rev != nil && rev.Val == slow.Val {
		slow = slow.Next
		rev = rev.Next
	}

	if rev != nil {
		return true
	}

	return false
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
}
