package main

import "fmt"

type MonotonicQueue struct {
	data []int
}

func (mq *MonotonicQueue) push(n int) {
	for len(mq.data) > 0 && mq.data[len(mq.data)-1] < n {
		mq.data = mq.data[:len(mq.data)-1]
	}
	mq.data = append(mq.data, n)
}

func (mq *MonotonicQueue) max() int {
	return mq.data[0]
}

func (mq *MonotonicQueue) pop(n int) {
	if len(mq.data) > 0 && mq.data[0] == n {
		mq.data = mq.data[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	window := &MonotonicQueue{data: []int{}}
	var res []int

	for i := 0; i < len(nums); i++ {
		if i < k-1 { // fill the first k-1 of the window
			window.push(nums[i])
		} else { // window slide forward
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
	}

	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
