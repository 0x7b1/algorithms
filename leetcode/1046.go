package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func SliceToMaxHeap(arr []int) IntHeap {
	var h IntHeap

	heap.Init(&h)
	for _, e := range arr {
		heap.Push(&h, e)
	}

	return h
}

func lastStoneWeight(stones []int) int {
	h := SliceToMaxHeap(stones)

	for len(h) > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)

		if a != b {
			heap.Push(&h, a-b)
		}
	}

	for len(h) > 0 {
		return heap.Pop(&h).(int)
	}

	return 0
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{5, 1, 8}))
	fmt.Println(lastStoneWeight([]int{8, 6, 2}))
}
