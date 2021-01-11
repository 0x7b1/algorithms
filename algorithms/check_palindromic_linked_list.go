package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func (l *ListNode) isPalindrome(head *ListNode) bool {
	l = head
	return l.traverse(head)
}

func (l *ListNode) traverse(right *ListNode) bool {
	if right == nil {
		return true
	}

	res := l.traverse(right.next)
	res = res && (right.val == l.val)
	l = l.next

	return res
}

// TODO: Needs refinement
// Time: O(n), Space: O(1)
func (l *ListNode) reverse(head *ListNode) *ListNode {
	pre := new(ListNode)
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}

	return pre
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}

	fmt.Println(list.isPalindrome(nil), list.reverse(nil))
}
