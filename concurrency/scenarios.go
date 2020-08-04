package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring1(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
}

func case_single_channel() {
	c := make(chan string)
	go boring1("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("Finished")
}

func boring2(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()

	return c
}

func case_chan_generator() {
	c := boring2("case 2")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("Finished")
}

func case_chan_generator_multiple() {
	joe := boring2("Joe")
	ann := boring2("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("Finished")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { c <- <-input1 }()
	go func() { c <- <-input2 }()

	return c
}

func case_chan_fan_in_multiplexing() {
	c := fanIn(boring2("Joe"), boring2("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Finished")
}

type Message struct {
	str  string
	wait chan bool
}

func case_restoring_sequencing() {
	// TODO
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// This implementation will wait for the sum of all the request to respond
func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// This implementation will wait for the last request to come without block
// But if the requests separately are actually slow then it's a problem
func Google2(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return
}

// This implementation will collect all the results that are within the time constraint
// However, will discard results that didn't finish on time, even if they have
// some work in progress
func Google2_1(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(time.Millisecond * 80)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

// This will process replicated search functions of the same query
// and will return the fastest one
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// This will process the searches in parallel and will also gather
// those who are out of time, keeping the results that are processed
func Google3(query string) (results []Result) {
	c := make(chan Result)
	Web1 := fakeSearch("web")
	Web2 := fakeSearch("web")
	Image1 := fakeSearch("image")
	Image2 := fakeSearch("image")
	Video1 := fakeSearch("video")
	Video2 := fakeSearch("video")

	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(time.Millisecond * 80)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timed o ut")
			return
		}
	}

	return
}

func caseFakeFramework() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	// Using Google v3
	results := Google3("golang")

	// Using First function on 2 replicas
	//results := First("golang",
	//	fakeSearch("replica 1"),
	//	fakeSearch("replica 2"),
	//)

	// Using Google v2.1
	//results := Google2_1("golang")

	// Using Google v2
	//results := Google2("golang")

	// Using Google v1
	//results := Google("golang")

	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func main() {
	//case_single_channel()
	//case_chan_generator()
	//case_chan_generator_multiple()
	//case_chan_fan_in_multiplexing()
	caseFakeFramework()
}
