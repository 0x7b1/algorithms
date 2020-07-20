package main

import "fmt"

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */
func findSolution2(customFunction func(int, int) int, z int) [][]int {
	var solutions [][]int

	for i := 1; i <= z; i++ {
		for j := i + 1; j <= z; j++ {
			if customFunction(i, j) == z {
				solutions = append(solutions, []int{i, j})
			}

			if customFunction(j, i) == z {
				solutions = append(solutions, []int{j, i})
			}
		}
	}

	for i := 1; i <= z; i++ {
		if customFunction(i, i) == z {
			solutions = append(solutions, []int{i, i})
		}
	}

	return solutions
}

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var solutions [][]int
	x, y := 1, 1

	for customFunction(x, y) < z {
		y++
	}

	for y > 0 {
		for y > 0 && customFunction(x, y) > z {
			y--
		}
		if y > 0 && customFunction(x, y) == z {
			solutions = append(solutions, []int{x, y})
		}
		x++
	}

	return solutions
}

func main() {
	fmt.Println(findSolution(func(x, y int) int { return x + y }, 5))
	fmt.Println(findSolution(func(x, y int) int { return x * y }, 5))
}
