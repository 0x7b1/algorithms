package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	i, j := len(num1) - 1, len(num2) - 1
	carry := 0
	res := ""

	for i >= 0 || j >= 0 || carry > 0 {
		sum := 0
		if i>= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j>= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		sum += carry
		carry = sum / 10
		sum = sum % 10
		res = string(sum + '0') + res
	}

	return res
}

func main() {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
	fmt.Println(addStrings("10", "20"))
}
