package leetcode

func minimumPossibleSum(n int, target int) int64 {
	if target/2 >= n {
		return int64((1 + n) * n / 2)
	}
	mid := target / 2
	rightLength := n - mid
	rightMax := target + rightLength - 1
	return int64((1+mid)*mid/2 + (target+rightMax)*rightLength/2)
}
