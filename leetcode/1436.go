package main

import "fmt"

func destCity(paths [][]string) string {
	dest := make(map[string]bool)

	for _, path := range paths {
		dest[path[0]] = true
	}

	for _, path := range paths {
		_, ok := dest[path[1]]
		if !ok {
			return path[1]
		}
	}

	return ""
}

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
}
