package main

import (
	"container/heap"
	"fmt"
)

type RowCounter struct {
	row     int
	counter int // a.k.a priority
	index   int
}

type PriorityQueue []*RowCounter

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].counter == pq[j].counter {
		return pq[i].row < pq[j].row
	} else {
		return pq[i].counter < pq[j].counter
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	rowCounter := x.(*RowCounter)
	rowCounter.index = len(*pq)
	*pq = append(*pq, rowCounter)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]

	return item
}

func countOnes(row []int) int {
	low, mid, hi := 0, 0, len(row)
	for low < hi {
		mid = low + (hi-low)/2
		if row[mid] == 1 {
			low = mid + 1
		} else {
			hi = mid
		}
	}

	return low
}

func kWeakestRows(mat [][]int, k int) []int {
	pq := make(PriorityQueue, len(mat))
	var weakestRowsIdx []int

	for i, row := range mat {
		pq[i] = &RowCounter{
			row:     i,
			counter: countOnes(row),
			index:   i,
		}

	}

	heap.Init(&pq)

	for k > 0 {
		counter := heap.Pop(&pq).(*RowCounter)
		weakestRowsIdx = append(weakestRowsIdx, counter.row)
		k--
	}

	return weakestRowsIdx
}

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 5))

	fmt.Println(kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}, 2))

	fmt.Println(kWeakestRows([][]int{
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
	}, 1))
}
