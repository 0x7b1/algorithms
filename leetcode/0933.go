package main

import "fmt"

type RecentCounter struct {
	t []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.t = append(this.t, t)

	idx := binarySearch(this.t, t - 3000)

	return len(this.t) - idx
}

func binarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] > target {
			hi = mid - 1
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return lo
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
	fmt.Println(obj.Ping(200))
}
