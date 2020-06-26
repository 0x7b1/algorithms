// -- MaxHeap --
// We use an array to represent a heap
// Root is at arr[0]
// Left child: arr[2 * i + 1]
// Right child: arr[2 * i + 2]
//
// Main operations:
// - Insert
// - Remove max element
// - Get Maximum element
package main

import "fmt"

type maxHeap struct {
	heapArray []int
	size      int
	maxsize   int
}

func newMaxHeap(maxsize int) *maxHeap {
	maxHeap := &maxHeap{
		heapArray: []int{},
		size:      0,
		maxsize:   maxsize,
	}

	return maxHeap
}

func (m *maxHeap) Insert(item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("Heap is full")
	}

	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)

	return nil
}

func (m *maxHeap) BuildMaxHeap() {
	for index := (m.size / 2) - 1; index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *maxHeap) downHeapify(index int) {
	// If is leaf node
	if index >= m.size/2 && index <= m.size {
		return
	}

	largest := index
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	if leftChildIndex < m.size && m.heapArray[leftChildIndex] > m.heapArray[largest] {
		largest = leftChildIndex
	}

	if rightChildIndex < m.size && m.heapArray[rightChildIndex] > m.heapArray[largest] {
		largest = rightChildIndex
	}

	if largest != index {
		m.swap(index, largest)
		m.downHeapify(largest)
	}
}

func (m *maxHeap) upHeapify(index int) {
	for m.heapArray[index] > m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *maxHeap) swap(first, second int) {
	m.heapArray[first], m.heapArray[second] = m.heapArray[second], m.heapArray[first]
}

func (m *maxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *maxHeap) Remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)

	return top
}

func main() {
	arr := []int{6, 5, 4, 7, 8, 3}
	maxHeap := newMaxHeap(len(arr))
	for _, e := range arr {
		maxHeap.Insert(e)
	}

	for range arr {
		fmt.Println("->", maxHeap.Remove())
	}
}
