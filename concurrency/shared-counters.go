package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func basic() {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	mu.Lock()
	fmt.Println(counter)
	mu.Unlock()
}

func counterTransfers() {
	alice, bob := 1000, 1000
	total := alice + bob
	var mu sync.Mutex

	go func() {
		for i := 0; i < 1000; i++ {
			// Locks protect invariance and shared data
			mu.Lock()
			alice--
			bob++
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice++
			bob--
			mu.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Println("leak not equal", alice, bob, total)
		}
		mu.Unlock()
	}
}

func voteCount() {
	rand.Seed(time.Now().UnixNano())
	count, finished := 0, 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			// voting
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			vote := rand.Int()
			mu.Lock()
			defer mu.Unlock()
			if vote%2 == 0 {
				count++
			}

			finished++
			cond.Broadcast()
		}()
	}

	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait()
	}

	if count >= 5 {
		fmt.Println("Received +5 votes")
	} else {
		fmt.Println("Lost")
	}

	mu.Unlock()
}

func main() {
	//basic()
	//counterTransfers()
	voteCount()
}
