package leetcode

func findMatrix(nums []int) [][]int {
	numMap := make(map[int]int, len(nums))
	for _, num := range nums {
		numMap[num]++
	}
	res := make([][]int, 0, len(nums))
	count := len(nums)
	for count > 0 {
		temp := make([]int, 0, count)
		for k := range numMap {
			if numMap[k] > 0 {
				temp = append(temp, k)
				numMap[k]--
				count--
			}
		}
		res = append(res, temp)
	}
	return res
}
