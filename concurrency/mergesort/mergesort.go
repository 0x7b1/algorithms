package mergesort

import (
	"math/rand"
	"sync"
	"time"
)

const max = 1 << 11

// https://medium.com/@teivah/parallel-merge-sort-in-go-fe14c1bc006

func random(n int) []int {
	s := make([]int, n)

	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}

	return s
}

func mergeInplace(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func mergeSortSequentialInplace(arr []int) {
	if len(arr) > 1 {
		middle := len(arr) / 2

		mergeSortSequentialInplace(arr[:middle])
		mergeSortSequentialInplace(arr[middle:])
		mergeInplace(arr, middle)
	}
}

func merge(a, b []int) []int {
	i, j := 0, 0
	var c []int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	for i < len(a) {
		c = append(c, a[i])
		i++
	}

	for j < len(b) {
		c = append(c, b[j])
		j++
	}

	return c
}

func mergeSortSequential(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	middle := len(a) / 2
	left := mergeSortSequential(a[:middle])
	right := mergeSortSequential(a[middle:])

	return merge(left, right)
}

func mergeSortParallel1(arr []int) {
	arrLen := len(arr)

	if arrLen > 1 {
		middle := arrLen / 2
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			mergeSortParallel1(arr[:middle])
		}()

		go func() {
			defer wg.Done()
			mergeSortParallel1(arr[middle:])
		}()

		wg.Wait()
		mergeInplace(arr, middle)
	}
}

func mergeSortParallel2(arr []int) {
	arrLen := len(arr)

	if arrLen > 1 {
		if arrLen <= max {
			mergeSortSequentialInplace(arr)
		} else {
			mergeSortParallel1(arr)
		}
	}
}

func mergeSortParallel3(arr []int) {
	arrLen := len(arr)

	if arrLen > 1 {
		if arrLen <= max {
			mergeSortSequentialInplace(arr)
		} else {
			middle := arrLen / 2
			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				mergeSortParallel3(arr[:middle])
			}()

			mergeSortParallel3(arr[middle:])

			wg.Wait()
			mergeInplace(arr, middle)
		}
	}
}

//func main() {
//	f, err := os.Create("trace.out")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	err = trace.Start(f)
//	if err != nil {
//		panic(err)
//	}
//	defer trace.Stop()
//
//	arr := []int{4, 3, 6, 1, 2}
//	mergeSortSequentialInplace(arr)
//	fmt.Println(arr)
//}
