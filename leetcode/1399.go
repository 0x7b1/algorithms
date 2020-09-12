package main

import "fmt"

func sumDigits(n int) (res int) {
	for n > 0 {
		res += n % 10
		n /= 10
	}

	return
}

func countLargestGroup(n int) int {
	groups := make(map[int]int)

	// group the numbers
	for i := 1; i <= n; i++ {
		groups[sumDigits(i)]++
	}

	// get the max length of the group
	var max int
	aux := make(map[int]int)
	for _, c := range groups {
		if c > max {
			max = c
		}
		aux[c]++
	}

	return aux[max]
}

func main() {
	fmt.Println(countLargestGroup(13))
	fmt.Println(countLargestGroup(2))
	fmt.Println(countLargestGroup(15))
	fmt.Println(countLargestGroup(24))
}
