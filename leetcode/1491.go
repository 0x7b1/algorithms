package main

import "fmt"

func average(salary []int) float64 {
	sum, min, max := 0, salary[0], salary[0]
	for _, n := range salary {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}
	fmt.Println(min, max)
	return float64(float64(sum - (max + min)) / float64(len(salary) - 2))
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}
