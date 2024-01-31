package leetcode

func dailyTemperatures(t []int) []int {
	n := len(t)
	stack := make([]int, 0, n)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && t[i] >= t[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}

func dailyTemperatures_2(temperatures []int) []int {
	n := len(temperatures)
	type pair struct {
		t   int
		idx int
	}
	stack := make([]pair, 0, n)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= stack[len(stack)-1].t {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1].idx - i
		}
		stack = append(stack, pair{t: temperatures[i], idx: i})
	}
	return res
}
