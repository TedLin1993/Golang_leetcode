package leetcode

func findLeastNumOfUniqueInts(arr []int, k int) int {
	n := len(arr)
	numCount := make(map[int]int, n)
	for _, v := range arr {
		numCount[v]++
	}
	counts := make([]int, k+1)
	for _, v := range numCount {
		if v <= k {
			counts[v]++
		}
	}
	res := len(numCount)
	for i := 1; i < len(counts) && k >= i; i++ {
		if k >= i*counts[i] {
			k -= i * counts[i]
			res -= counts[i]
		} else {
			return res - k/i
		}
	}
	return res
}
