package leetcode

//divide to 2 part
//left part: 1 + 2 + ... + k/2
//right part: k + k+1 + ... + k+(n-k/2)-1
func minimumSum(n int, k int) int {
	if k/2 >= n {
		return (1 + n) * n / 2
	}
	leftLen := k / 2
	rightLen := n - leftLen
	leftSum := (1 + leftLen) * leftLen / 2
	rightSum := (k + k + rightLen - 1) * rightLen / 2
	return leftSum + rightSum
}
