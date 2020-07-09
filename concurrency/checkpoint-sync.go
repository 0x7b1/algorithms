package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	partList := []string{"A", "B", "C", "D"}
	nAssemblies := 3
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nAssemblies; i++ {
		fmt.Println("Begin assembly", i)
		wg.Add(len(partList))

		for _, part := range partList {
			go func(part string, wg *sync.WaitGroup) {
				fmt.Println("begin part", part)
				time.Sleep(time.Millisecond * 500)
				fmt.Println("end part", part)
				wg.Done()
			}(part, &wg)
		}

		wg.Wait()
		fmt.Println("Assembly", i, "complete")
	}
}
