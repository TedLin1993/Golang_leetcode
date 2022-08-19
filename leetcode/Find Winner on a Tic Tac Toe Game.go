package leetcode

func tictactoe(moves [][]int) string {
	matrix := make([][]string, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]string, 3)
	}
	for idx, move := range moves {
		player := "A"
		if idx%2 == 1 {
			player = "B"
		}
		matrix[move[0]][move[1]] = player
	}

	//check row
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] != "" && matrix[i][0] == matrix[i][1] && matrix[i][1] == matrix[i][2] {
			return matrix[i][0]
		}
	}
	//check column
	for i := 0; i < len(matrix); i++ {
		if matrix[0][i] != "" && matrix[0][i] == matrix[1][i] && matrix[1][i] == matrix[2][i] {
			return matrix[0][i]
		}
	}
	//check diagonal
	if matrix[0][0] != "" && matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] {
		return matrix[0][0]
	}
	if matrix[2][0] != "" && matrix[2][0] == matrix[1][1] && matrix[1][1] == matrix[0][2] {
		return matrix[2][0]
	}
	
	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}
