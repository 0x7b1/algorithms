package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq IntHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *IntHeap) Pop() interface{} {
	old := *pq
	n := len(*pq) - 1
	last := old[n]
	*pq = old[:n]
	return last
}

type KthLargest struct {
	k  int
	pq *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap{}
	kthLargest := KthLargest{k, &h}
	heap.Init(&h)

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.pq, val)
	if this.pq.Len() > this.k {
		heap.Pop(this.pq)
	}

	return (*this.pq)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func streamNumbers() {

}

func main() {
	//streamNumbers()

	k := 3
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, arr)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))
}
