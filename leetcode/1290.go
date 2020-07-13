package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	dec := 0
	for head != nil {
		dec = dec*2 + head.Val
		head = head.Next
	}

	return dec
}

func main() {
	fmt.Println(getDecimalValue(&ListNode{1, &ListNode{0, &ListNode{1, nil}}}))
}
