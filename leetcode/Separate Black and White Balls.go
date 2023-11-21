package leetcode

func minimumSteps(s string) int64 {
	left := -1
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			res += i - left - 1
			left++
		}
	}
	return int64(res)
}
