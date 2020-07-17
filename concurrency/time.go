package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func periodic() {
	for {
		fmt.Println("TICK")
		time.Sleep(time.Second)
	}
}

func sleep() {
	time.Sleep(time.Second)
	fmt.Println("Started")
	go periodic()
	time.Sleep(time.Second * 5)
}

func periodicCancel() {
	for {
		fmt.Println("Ticking")
		time.Sleep(time.Second)
		mu.Lock()
		if done {
			fmt.Println("killing")
			return
		}
		mu.Unlock()
	}
}

func sleepCancel() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("starting")
	go periodicCancel()
	time.Sleep(time.Second * 5)
	mu.Lock()
	done = true
	mu.Unlock()
	time.Sleep(time.Second * 3)
}

func main() {
	//sleep()
	sleepCancel()
}
