package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	res := make([]bool, len(candies))

	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	for i, c := range candies {
		if c+extraCandies >= max {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}

func main() {
	a := []int{2, 3, 5, 1, 3}
	fmt.Println(kidsWithCandies(a, 3))
}
