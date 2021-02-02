package main

import "fmt"

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	res := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mult := int((num1[i] - '0') * (num2[j] - '0'))
			p1 := i + j
			p2 := i + j + 1
			sum := mult + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	i := 0
	for ;i < len(res) && res[i] == 0; i++ {}

	str := ""
	for ; i <len(res); i++ {
		str += string('0' + res[i])
	}

	if len(str) == 0 {
		return "0"
	}

	return str
}

func main() {
	fmt.Println(multiply("123", "456"))
}
