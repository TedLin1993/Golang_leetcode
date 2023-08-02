package leetcode

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	var dfs func(comb []int, num int)
	dfs = func(comb []int, num int) {
		if len(comb) == k {
			temp := make([]int, k)
			copy(temp, comb)
			res = append(res, temp)
			return
		}
		for i := num; i <= n; i++ {
			dfs(append(comb, i), i+1)
		}
	}
	dfs(make([]int, 0, k), 1)
	return res
}

func Test_combine() {
	fmt.Println(combine(5, 4)) //[[1 2 3 4] [1 2 3 5] [1 2 4 5] [1 3 4 5] [2 3 4 5]]
}
