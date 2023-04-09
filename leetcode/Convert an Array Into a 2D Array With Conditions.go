package leetcode

func findMatrix(nums []int) [][]int {
	numMap := make(map[int]int, len(nums))
	rowCount := 0
	for _, num := range nums {
		numMap[num]++
		if numMap[num] > rowCount {
			rowCount = numMap[num]
		}
	}
	res := make([][]int, rowCount)
	for k := range numMap {
		for i := 0; i < numMap[k]; i++ {
			res[i] = append(res[i], k)
		}
	}
	return res
}
