package leetcode

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	minLen := min(len(s1), min(len(s2), len(s3)))
	commonLen := 0
	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] || s2[i] != s3[i] {
			break
		}
		commonLen = i + 1
	}
	if commonLen == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - 3*commonLen
}
