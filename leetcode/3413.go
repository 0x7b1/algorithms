package main

import "fmt"

// TODO
type WordDictionary struct {
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {

}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return false
}

func main() {
	obj := Constructor()
	obj.AddWord("abc")
	param_2 := obj.Search(".bc")
	fmt.Println(param_2)
}
