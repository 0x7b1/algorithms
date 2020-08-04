package main

import "fmt"

type MyHashSet struct {

}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{}
}


func (this *MyHashSet) Add(key int)  {

}


func (this *MyHashSet) Remove(key int)  {

}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
	obj := Constructor()
	obj.Add(2)
	obj.Remove(3)
	param_3 := obj.Contains(2)

	fmt.Println(param_3)
}
