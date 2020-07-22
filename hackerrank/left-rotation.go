package main

import "fmt"

func rotateLeft(n, r int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	copy(arr, append(arr[r:], arr[0:r]...))

	fmt.Println(arr)
}

func main() {
	//rotateLeft(5, 3)
	rotateLeft(5, 4)
}
