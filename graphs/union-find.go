package main

import "fmt"

type UF struct {
	counter int
	parent  []int

	// size keeps track of the nr of nodes in each tree
	size []int
}

// union connects two components
// Time: O(1)
func (u *UF) union(p, q int) {
	rootP, rootQ := u.find(p), u.find(q)

	if rootP == rootQ {
		return
	}

	// Merge two trees into one
	// The small tree is more balanced
	if u.size[rootP] > u.size[rootQ] {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	} else {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	}

	u.counter--
}

// find returns the root node of a node
// Time: O(1)
func (u *UF) find(x int) int {
	for u.parent[x] != x {
		// Path compression
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}

	return x
}

// connected checks if two components have the same root
// Time: O(1)
func (u *UF) connected(p, q int) bool {
	rootP, rootQ := u.find(p), u.find(q)

	return rootP == rootQ
}

// count returns the number of connected components
func (u *UF) count() int {
	return u.counter
}

// Time: O(n)
func NewUF(n int) *UF {
	uf := new(UF)

	// Disconnected at first
	uf.counter = n

	// Parent node points to itself
	uf.parent = make([]int, n)
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func main() {
	uf := NewUF(10)
	uf.union(0, 2)
	uf.union(3, 2)
	uf.union(4, 9)
	fmt.Println(uf.count())
}
