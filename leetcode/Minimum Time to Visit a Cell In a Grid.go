package leetcode

import (
	"container/heap"
	"fmt"
)

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	minTimeHeap := MinTimeHeap{{Time: 0, Cell: []int{0, 0}}}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(minTimeHeap) > 0 {
		cellTime := heap.Pop(&minTimeHeap).(CellTime)
		if cellTime.Cell[0] == m-1 && cellTime.Cell[1] == n-1 {
			return cellTime.Time
		}

		for _, dir := range directions {
			r, c := cellTime.Cell[0]+dir[0], cellTime.Cell[1]+dir[1]
			if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] {
				continue
			}
			wait := 0
			if grid[r][c] > cellTime.Time+1 && (grid[r][c]-cellTime.Time)%2 == 0 {
				wait = 1
			}
			time := max(cellTime.Time+1, grid[r][c]+wait)
			heap.Push(&minTimeHeap, CellTime{Time: time, Cell: []int{r, c}})
			visited[r][c] = true
		}
	}
	return -1
}

type CellTime struct {
	Time int
	Cell []int
}

type MinTimeHeap []CellTime

func (h MinTimeHeap) Len() int           { return len(h) }
func (h MinTimeHeap) Less(i, j int) bool { return h[i].Time < h[j].Time }
func (h MinTimeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinTimeHeap) Push(x interface{}) {
	*h = append(*h, x.(CellTime))
}
func (h *MinTimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Test_minimumTime() {
	fmt.Println(minimumTime([][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}})) //7
}
