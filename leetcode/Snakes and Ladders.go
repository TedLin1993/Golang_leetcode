package leetcode

// n = len(board)
// label's row = n - (label/n) - 1
// label's column = label/n %2 == 0 ? label%n: n- label%n

func snakesAndLadders(board [][]int) int {
	n := len(board)
	maxLabel := n * n
	res := 0
	visited := make([]bool, maxLabel+1)
	visited[1] = true
	queue := make([]int, 0, maxLabel)
	queue = append(queue, 1)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			label := queue[i]
			for j := 1; j <= 6; j++ {
				nextLabel := label + j
				if nextLabel > maxLabel {
					continue
				}
				r, c := getIndexBylabel(nextLabel, n)
				if board[r][c] != -1 {
					nextLabel = board[r][c]
				}
				if nextLabel == maxLabel {
					return res + 1
				}
				if !visited[nextLabel] {
					visited[nextLabel] = true
					queue = append(queue, nextLabel)
				}
			}
		}
		res++
		queue = queue[length:]
	}

	return -1
}

func getIndexBylabel(label, n int) (row, col int) {
	idx := label - 1
	row = n - (idx / n) - 1
	if (idx/n)%2 == 0 {
		col = idx % n
	} else {
		col = n - idx%n - 1
	}
	return
}
