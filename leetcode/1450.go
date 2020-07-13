package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	size, count := len(startTime), 0
	for i := 0; i < size; i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
}
