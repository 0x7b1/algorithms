package main

import "fmt"

// This could be improved with a max heap
func repeatedNTimes(A []int) int {
	counters := make(map[int]int)
	for _, n := range A {
		counters[n]++
	}

	max := A[0]

	for k, v := range counters {
		if v > counters[max] {
			max = k
		}
	}

	return max

}

func main() {
	fmt.Println(repeatedNTimes([]int{2,1,2,5,3,2}))
}
