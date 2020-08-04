package main

import "fmt"

// Lame and shameful implementation of a HashSet
type MyHashSet struct {
	contents []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	contents := []int{}
	return MyHashSet{
		contents: contents,
	}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.contents = append(this.contents, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	for i := 0; i < len(this.contents); i++ {
		if this.contents[i] == key {
			this.contents = append(this.contents[:i], this.contents[i+1:]...)
			break
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for _, n := range this.contents {
		if n == key {
			return true
		}
	}

	return false
}

func main() {
	obj := Constructor()
	obj.Add(2)
	obj.Remove(3)
	obj.Add(20)
	obj.Add(44)
	obj.Add(44)
	param_3 := obj.Contains(44)

	fmt.Println(param_3)
}
