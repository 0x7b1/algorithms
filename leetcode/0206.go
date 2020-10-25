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

func (ln *ListNode) String() string {
	return fmt.Sprintf("%v -> %v", ln.Val, ln.Next)
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	invList1 := reverseList(list1)
	fmt.Println(invList1)
}
