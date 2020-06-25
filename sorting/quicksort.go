package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivotIndex := rnd.Intn(len(a))

	a[pivotIndex], a[right] = a[pivotIndex], a[right]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[right], a[left] = a[left], a[right]
	quickSort(a[:left])
	quickSort(a[left+1:])
}

func main() {
	arr := []int{5, 4, 3, 22, 2, 1}
	quickSort(arr)
	fmt.Println(arr)
}
