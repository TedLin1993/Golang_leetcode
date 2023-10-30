package leetcode

func minIncrementOperations(nums []int, k int) int64 {
	dp0, dp1, dp2 := 0, 0, 0
	for _, num := range nums {
		incre := max(k-num, 0)
		dp0, dp1, dp2 = min(dp0, min(dp1, dp2))+incre, dp0, dp1
	}
	return int64(min(dp0, min(dp1, dp2)))
}
