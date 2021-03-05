package main

import "fmt"

type MyQueue struct {
	input, output *[]int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		input:  new([]int),
		output: new([]int),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	*this.input = append(*this.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v := this.Peek()
	*this.output = (*this.output)[:len(*this.output)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*this.output) == 0 {
		for len(*this.input) > 0 {
			*this.output = append(*this.output, (*this.input)[len(*this.input)-1])
			*this.input = (*this.input)[:len(*this.input)-1]
		}
	}

	return (*this.output)[len(*this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*this.input) == 0 && len(*this.output) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
