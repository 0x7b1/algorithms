package main

import (
	"fmt"
	"sync"
	"time"
)

func sendRPC(i int) {
	fmt.Println("->", i)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Millisecond * 500)
			sendRPC(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
