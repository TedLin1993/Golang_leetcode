package leetcode

func findChampion(grid [][]int) int {
	res := -1
	max := 0
	for i, row := range grid {
		temp := 0
		for _, v := range row {
			if v == 1 {
				temp++
			}
		}
		if temp > max {
			max = temp
			res = i
		}
	}
	return res
}
