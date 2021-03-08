package main

import "fmt"

func convertToTitle(columnNumber int) string {
	var res []byte

	for columnNumber > 0 {
		res = append(res, 'A'+byte((columnNumber-1)%26))
		columnNumber = (columnNumber - 1) / 26
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(2147483647))
}
