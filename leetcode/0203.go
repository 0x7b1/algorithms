package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	node := head

	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	if head.Val == val {
		head = head.Next
	}

	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%v -> %v", l.Val, l.Next)
}

func main() {
	list := &ListNode{1,
		&ListNode{2,
			&ListNode{6,
				&ListNode{3,
					&ListNode{6, nil}}}}}

	fmt.Println(removeElements(list, 6))

	list2 := &ListNode{7,
				&ListNode{7,
			&ListNode{7, nil}}}

	fmt.Println(removeElements(list2, 7))
}
