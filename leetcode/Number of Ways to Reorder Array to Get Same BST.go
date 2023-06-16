package leetcode

import "fmt"

func numOfWays(nums []int) int {
	modulo := int(1e9 + 7)
	table := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		table[i] = make([]int, len(nums))
		table[i][0] = 1
		for j := 1; j <= i; j++ {
			table[i][j] = (table[i-1][j-1] + table[i-1][j]) % modulo
		}
	}
	var dfs func(arr []int) int
	dfs = func(arr []int) int {
		n := len(arr)
		if n <= 2 {
			return 1
		}
		left, right := make([]int, 0, n), make([]int, 0, n)
		firstNode := arr[0]
		for _, num := range arr[1:] {
			if num < firstNode {
				left = append(left, num)
			} else {
				right = append(right, num)
			}
		}
		m := len(left)
		res := table[n-1][m]
		res = res * dfs(left) % modulo
		res = res * dfs(right) % modulo
		return res
	}
	return dfs(nums) - 1
}

func Test_numOfWays() {
	// fmt.Println(numOfWays([]int{2, 1, 3}))                                                                                                                               //1
	// fmt.Println(numOfWays([]int{3, 4, 5, 1, 2}))                                                                                                                         //5
	// fmt.Println(numOfWays([]int{1, 2, 3}))                                                                                                                               //0
	// fmt.Println(numOfWays([]int{3, 1, 2, 5, 4, 6}))                                                                                                                      //19
	// fmt.Println(numOfWays([]int{31, 23, 14, 24, 15, 12, 25, 28, 5, 35, 17, 6, 9, 11, 1, 27, 18, 20, 2, 3, 33, 10, 13, 4, 7, 36, 32, 29, 8, 30, 26, 19, 34, 22, 21, 16})) //936157466
	fmt.Println(numOfWays([]int{10, 23, 12, 18, 4, 29, 2, 8, 41, 31, 25, 21, 14, 35, 26, 5, 19, 43, 22, 37, 9, 20, 44, 28, 1, 39, 30, 38, 36, 6, 13, 16, 27, 17, 34, 7, 15, 3, 11, 24, 42, 33, 40, 32})) //182440977
}
