package leetcode

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if checkExistDfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func checkExistDfs(board [][]byte, row int, column int, word string, wordIdx int) bool {
	if board[row][column] != word[wordIdx] {
		return false
	}
	if wordIdx == len(word)-1 {
		return true
	}
	char := board[row][column]
	board[row][column] = '$'

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := column + direction[1]
		if nextRow >= 0 && nextRow < len(board) && nextCol >= 0 && nextCol < len(board[0]) {
			if checkExistDfs(board, nextRow, nextCol, word, wordIdx+1) {
				return true
			}
		}
	}

	board[row][column] = char
	return false
}

func Test_wordSearch() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")) //true
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))   //false
	fmt.Println(exist([][]byte{{'A'}}, "A"))                                                                 //true
}
