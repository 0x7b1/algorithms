package main

import "fmt"

func buildArray(target []int, n int) []string {
	var ops []string
	curr := 1

	for _, t := range target {
		for i := curr; i < t; i++ {
			ops = append(ops, "Push")
			ops = append(ops, "Pop")
		}
		ops = append(ops, "Push")
		curr = t + 1
	}

	return ops
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{2, 3, 4}, 4))
	fmt.Println(buildArray([]int{1, 2}, 4))
}
