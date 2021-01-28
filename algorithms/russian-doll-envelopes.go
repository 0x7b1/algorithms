package main

import (
	"fmt"
	"sort"
)

type HeightWidthPair []int

func (h HeightWidthPair) Len() int {

}

func (h HeightWidthPair) Less(i, j int) bool {

}

func (h * HeightWidthPair) Swap(i, j int)  {

}

func maxEnvelopes(envelopes []HeightWidthPair) []HeightWidthPair {
	//n := len(envelopes)
	sort.Sort(envelopes)

}

func main() {
	fmt.Println(maxEnvelopes([]HeightWidthPair{
		{2, 20},
		{1, 10},
		{4, 40},
	}))
}
