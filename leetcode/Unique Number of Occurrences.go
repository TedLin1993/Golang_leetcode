package leetcode

func uniqueOccurrences(arr []int) bool {
	arrMap := make(map[int]int, len(arr))
	for _, v := range arr {
		arrMap[v]++
	}
	countVisitMap := make(map[int]bool, len(arrMap))
	for _, count := range arrMap {
		if countVisitMap[count] {
			return false
		}
		countVisitMap[count] = true
	}
	return true
}
