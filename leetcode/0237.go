package main

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next != nil {
		node.Val, node.Next = node.Next.Val, node.Next.Next
	}
}

func main() {
	lst := &ListNode{}
	deleteNode(lst)
}
