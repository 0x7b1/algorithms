package main

import "fmt"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	chars := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	result := ""

	for num != 0 {
		result = string(chars[(num & 15)]) + result
		num = int(uint32(num) >> 4)
	}

	return result
}

func main() {
	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
}
