package main

import "fmt"

type MyStack struct {
	q1, q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{q1: make([]int, 0), q2: make([]int, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.q1)
	for i := 0; i < length-1; i++ {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}

	topElement := this.q1[0]
	this.q1 = this.q2
	this.q2 = nil

	return topElement
}

/** Get the top element. */
func (this *MyStack) Top() int {
	topElement := this.Pop()
	this.q1 = append(this.q1, topElement)
	return topElement
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.q1) == 0 {
		return true
	}

	return false
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.Empty())

}
