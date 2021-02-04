package main

import "fmt"

func searchWord(board [][]byte, height, width int, word string, x, y, i int) bool {
	if len(word) == i {
		return true
	}

	if y < 0 || y >= height || x < 0 || x >= width {
		return false
	}

	if board[y][x] != word[i] {
		return false
	}

	board[y][x] ^= 0xFF

	exist := searchWord(board, height, width, word, x-1, y, i+1) ||
		searchWord(board, height, width, word, x+1, y, i+1) ||
		searchWord(board, height, width, word, x, y-1, i+1) ||
		searchWord(board, height, width, word, x, y+1, i+1)

	board[y][x] ^= 0xFF

	return exist
}

func exist(board [][]byte, word string) bool {
	height, width := len(board), len(board[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if searchWord(board, height, width, word, x, y, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}
