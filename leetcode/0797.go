package main

import "fmt"

func dfs(current int, graph *[][]int) [][]int {
	paths := [][]int{}

	if current == len(*graph)-1 {
		return [][]int{{current}}
	}

	for _, node := range (*graph)[current] {
		t := dfs(node, graph)
		for _, path := range t {
			paths = append(paths, append([]int{current}, path...))
		}
	}

	return paths
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return dfs(0, &graph)
}

func main() {
	g1 := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	g2 := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}

	fmt.Println(allPathsSourceTarget(g1))
	fmt.Println(allPathsSourceTarget(g2))
}
