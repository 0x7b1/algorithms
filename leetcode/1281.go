package main

import "fmt"

func subtractProductAndSum(n int) int {
	product, sum := 1, 0

	for n > 0 {
		product *= n % 10
		sum += n % 10
		n /= 10
	}

	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}
