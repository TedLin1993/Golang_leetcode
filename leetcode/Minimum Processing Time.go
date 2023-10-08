package leetcode

func differenceOfSums(n int, m int) int {
	var k = n / m
	return n*(n+1)/2 - k*(k+1)*m
}
