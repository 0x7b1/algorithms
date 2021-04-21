package main

import "fmt"

func strWithout3a3b(a int, b int) string {
	charA, charB := "a", "b"

	if b > a {
		a, b = b, a
		charA, charB = charB, charA
	}

	res := ""
	for a > 0 {
		a--
		res += charA
		if a > b {
			res += charA
			a--
		}

		if b > 0 {
			b--
			res += charB
		}
	}

	return res
}

func main() {
	fmt.Println(strWithout3a3b(1, 2))
	fmt.Println(strWithout3a3b(4, 1))
}
