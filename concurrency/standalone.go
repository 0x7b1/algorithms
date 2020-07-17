package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func countOnesInRow(arr []int) int {
	lo, mi, hi := 0, 0, len(arr)

	for lo < hi {
		mi = lo + (hi - lo) / 2
		if arr[mi] == 1 {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return lo
}

func concurrentRowCountMutex(arr [][]int) []int {
	rowCounter := make([]int, len(arr))
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(arr))
	for i, row := range arr {
		go func(i int, row []int) {
			count := countOnesInRow(row)
			mu.Lock()
			rowCounter[i] = count
			mu.Unlock()
			wg.Done()
		}(i, row)
	}

	wg.Wait()

	return rowCounter
}

func concurrentRowCountChannels(arr [][]int) []int {
	type kv struct{ row, count int }

	countChan := make(chan *kv, len(arr))

	for i, row := range arr {
		go func(i int, row []int) {
			count := countOnesInRow(row)
			countChan <- &kv{i, count}
		}(i, row)
	}

	//close(countChan) // TODO: Fix closing behavior

	collector := make([]int, len(arr))

	time.Sleep(time.Millisecond * 500)

	fmt.Println(<-countChan)
	fmt.Println(<-countChan)

	// TODO: Fix this out-of-order channel message consumption
	//for range countChan {
	//	fmt.Println("donando")
	//	v := <-countChan
	//	collector[v.row] = v.count
	//}

	return collector
}

func sequentialRowCount(arr [][]int) []int {
	rowCount := make([]int, len(arr))
	var counter int

	for i, row := range arr {
		for _, e := range row {
			counter += e
		}

		rowCount[i] = counter
		counter = 0
	}

	return rowCount
}

func rowMatrixGenerator() [][]int {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10) + 1
	m := rand.Intn(10) + 1
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		row := make([]int, m)
		countOnes := rand.Intn(m) + 1
		for j := 0; j < countOnes; j++ {
			row[j] = 1
		}

		matrix[i] = row
	}

	return matrix
}

type Matrix [][]int

func (mt Matrix) String() string {
	n, m := len(mt), len(mt[0])
	str := fmt.Sprintf("%d x %d\n", n, m)

	str += "[\n"
	for _, row := range mt {
		str += fmt.Sprintf("  %v,\n", row)
	}
	str += "]\n"

	return str
}

func main() {
	//arr := [][]int{
	//	{1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	//	{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	//	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//}

	arr := rowMatrixGenerator()
	fmt.Println(Matrix(arr))

	fmt.Println(sequentialRowCount(arr))
	fmt.Println(concurrentRowCountMutex(arr))
	//fmt.Println(concurrentRowCountChannels(arr))
}
