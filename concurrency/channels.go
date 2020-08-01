package main

import (
	"fmt"
	"math/rand"
	"time"
)

func unbuffered() {
	c := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		<-c
	}()

	start := time.Now()
	c <- true
	fmt.Println("took", time.Since(start).Seconds())
}

func doWork(c chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		c <- rand.Int()
	}
}

func producerConsumer() {
	c := make(chan int)

	for i := 0; i < 4; i++ {
		go doWork(c)
	}

	for {
		v := <-c
		fmt.Println(v)
	}
}

func main() {
	//unbuffered()
	//producerConsumer()
	fmt.Println("Hello World amigos del mas alla")
}
