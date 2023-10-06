package leetcode

func maximumSumOfHeights_2866(maxHeights []int) int64 {
	n := len(maxHeights)
	stack := []int{n}
	suf := make([]int, n+1)
	cur := 0
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 1 && maxHeights[i] <= maxHeights[stack[len(stack)-1]] {
			cur -= maxHeights[stack[len(stack)-1]] * (stack[len(stack)-2] - stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		cur += maxHeights[i] * (stack[len(stack)-1] - i)
		stack = append(stack, i)
		suf[i] = cur
	}

	stack = []int{-1}
	pre := make([]int, n)
	cur = 0
	for i := 0; i < n; i++ {
		for len(stack) > 1 && maxHeights[i] <= maxHeights[stack[len(stack)-1]] {
			cur -= maxHeights[stack[len(stack)-1]] * (stack[len(stack)-1] - stack[len(stack)-2])
			stack = stack[:len(stack)-1]
		}
		cur += maxHeights[i] * (i - stack[len(stack)-1])
		stack = append(stack, i)
		pre[i] = cur
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, pre[i]+suf[i+1])
	}
	return int64(res)
}
