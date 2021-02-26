package main

import "fmt"

type MinStack struct {
	min   []int
	arr   []int
	index int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: make([]int, 0),
		arr: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
	if this.index == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}

	this.index++
}

func (this *MinStack) Pop() {
	this.index--
	this.arr = this.arr[:this.index]
	this.min = this.min[:this.index]
}

func (this *MinStack) Top() int {
	return this.arr[this.index - 1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.index - 1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
