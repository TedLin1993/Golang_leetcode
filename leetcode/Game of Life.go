package leetcode

import "fmt"

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			sum := 0
			for k := i - 1; k <= i+1; k++ {
				for r := j - 1; r <= j+1; r++ {
					if k < 0 || k >= len(board) || r < 0 || r >= len(board[0]) || (k == i && r == j) {
						continue
					}
					//origin is 1
					if board[k][r] == 1 || board[k][r] == -1 {
						sum++
					}
				}
			}
			val := board[i][j]
			// 1 become 1 or -1; 0 become 0 or 2
			if val == 1 {
				if sum < 2 || sum > 3 {
					board[i][j] = -1
				} else {
					board[i][j] = 1
				}
			} else if val == 0 {
				if sum == 3 {
					board[i][j] = 2
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func TestgameOfLife() {
	var val [][]int
	// val = [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// gameOfLife(val)
	// fmt.Println(val) //[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
	val = [][]int{{1, 1}, {0, 1}}
	gameOfLife(val)
	fmt.Println(val) //[[1,1],[1,1]]
}
