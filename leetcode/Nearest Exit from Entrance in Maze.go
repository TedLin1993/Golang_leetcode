package leetcode

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	visit := make([][]bool, len(maze))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(maze[0]))
	}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{entrance}
	visit[entrance[0]][entrance[1]] = true
	res := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if res != 0 && isBorder(queue[i], len(maze), len(maze[0])) {
				return res
			}
			x, y := queue[i][0], queue[i][1]
			for _, direction := range directions {
				next_x := x + direction[0]
				next_y := y + direction[1]
				if isValidPosition([]int{next_x, next_y}, maze, visit) {
					queue = append(queue, []int{next_x, next_y})
					visit[next_x][next_y] = true
				}
			}
		}
		res++
		queue = queue[length:]
	}
	return -1
}

func isBorder(position []int, row int, column int) bool {
	return position[0] == 0 || position[0] == row-1 || position[1] == 0 || position[1] == column-1
}

func isValidPosition(position []int, maze [][]byte, visit [][]bool) bool {
	x, y := position[0], position[1]
	if x >= 0 && x < len(maze) &&
		y >= 0 && y < len(maze[0]) &&
		!visit[x][y] && maze[x][y] == '.' {
		return true
	}
	return false
}

func Test_nearestExit() {
	fmt.Println(nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2})) //1
	fmt.Println(nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))                //2
}
