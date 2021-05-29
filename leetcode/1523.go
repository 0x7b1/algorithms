package main

import "fmt"

func countOdds(low int, high int) int {
	return (high + 1) / 2 - low / 2
}

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
}
