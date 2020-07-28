package main

import (
	"bytes"
	"fmt"
	"sort"
)

type pair struct {
	b     byte
	count int
}

func sortString(s string) string {
	var res bytes.Buffer
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	counts := []pair{}
	for i := 0; i < len(bs); {
		count := 1
		j := i + 1
		for ; j < len(bs); j++ {
			if bs[i] == bs[j] {
				count++
			} else {
				break
			}
		}

		counts = append(counts, pair{bs[i], count})
		i = j
	}

	finished := false
	for !finished {
		finished = true
		for i := range counts {
			if counts[i].count != 0 {
				finished = false
				res.WriteByte(counts[i].b)
				counts[i].count--
			}
		}
	}

	for i := len(counts) - 1; i >= 0; i-- {
		if counts[i].count != 0 {
			finished = false
			res.WriteByte(counts[i].b)
			counts[i].count--
		}
	}

	return res.String()
}

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
	fmt.Println(sortString("rat"))
}
