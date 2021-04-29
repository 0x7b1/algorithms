package main

import "fmt"

func fizzBuzz(n int) []string {
	res := []string{}

	for i := 1; i <= n; i++ {
		if i %3 == 0 && i % 5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i % 5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%v", i))
		}
	}

	return res
}

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}
