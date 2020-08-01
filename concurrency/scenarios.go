package main

import (
	"fmt"
	"os"
	"time"
)

func boring() <-chan int {
	c := make(chan int)

	for i := 0; ; i++ {
		c <- i
		time.Sleep(time.Millisecond * 300)
	}

	return c
}

func main() {
	fmt.Println("WORKSdwa cuba", os.Args)
}
