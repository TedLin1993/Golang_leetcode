package leetcode

import "fmt"

func isValidSudoku(board [][]byte) bool {
	//row and column
	for i := 0; i < len(board); i++ {
		rowMap := make(map[byte]bool)
		columnMap := make(map[byte]bool)
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && rowMap[board[i][j]] {
				return false
			}
			rowMap[board[i][j]] = true

			if board[j][i] != '.' && columnMap[board[j][i]] {
				return false
			}
			columnMap[board[j][i]] = true
		}
	}

	//sub-box
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			numMap := make(map[byte]bool)
			for k := i; k < i+3; k++ {
				for r := j; r < j+3; r++ {
					if board[k][r] != '.' && numMap[board[k][r]] {
						return false
					}
					numMap[board[k][r]] = true
				}
			}
		}
	}

	return true
}

func TestisValidSudoku() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}})) //true
}
