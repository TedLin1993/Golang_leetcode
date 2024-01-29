package leetcode

func countKeyChanges(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1]+32 && s[i] != s[i-1]-32 && s[i] != s[i-1] {
			res++
		}
	}
	return res
}
