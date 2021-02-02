package main

import "fmt"

func removeDuplicatesArray(nums []int) int {
	slow, fast := 0, 1
	n := len(nums)

	if n == 0 {
		return 0
	}

	for fast < n {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}

	return slow + 1
}

type Node struct {
	Val  int
	Next *Node
}

func removeDuplicatesLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next

	for fast != nil {
		if slow.Val != fast.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}

		fast = fast.Next
	}

	slow.Next = nil

	return head
}

func main() {
	a1 := []int{1, 1, 2}
	fmt.Println(removeDuplicatesArray(a1), a1)

	a2 := &Node{1, &Node{1, &Node{2, nil}}}
	fmt.Println(removeDuplicatesLinkedList(a2))
}
