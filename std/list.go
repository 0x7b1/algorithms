package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(10)
	intList.PushBack(9)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
