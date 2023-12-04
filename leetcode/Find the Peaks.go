package leetcode

func findPeaks(mountain []int) []int {
	n := len(mountain)
	res := make([]int, 0, n)
	for i := 1; i < n-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}
	return res
}
