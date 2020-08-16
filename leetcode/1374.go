package main

import "fmt"

func generateTheString(n int) string {
	out := ""

	if n % 2 == 0 {
		out = "b"
		n--
	}

	for i := 0; i < n; i++ {
		out += "a"
	}

	return out
}

func main() {
	fmt.Println(generateTheString(6))
}
