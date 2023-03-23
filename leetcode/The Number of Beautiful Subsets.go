package leetcode

import "fmt"

func beautifulSubsets(nums []int, k int) int {
	var dfs func(idx int, containsMap map[int]int) int
	dfs = func(idx int, containsMap map[int]int) int {
		if containsMap[nums[idx]+k] > 0 || containsMap[nums[idx]-k] > 0 {
			return 0
		}
		res := 1
		containsMap[nums[idx]]++
		for i := idx + 1; i < len(nums); i++ {
			res += dfs(i, containsMap)
		}
		containsMap[nums[idx]]--
		return res
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		containsMap := make(map[int]int, len(nums)-i)
		res += dfs(i, containsMap)
	}
	return res
}

func Test_beautifulSubsets() {
	fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2))           //4
	fmt.Println(beautifulSubsets([]int{2, 4, 6, 7}, 2))        //9
	fmt.Println(beautifulSubsets([]int{4, 2, 5, 9, 10, 3}, 1)) //23
	fmt.Println(beautifulSubsets([]int{1, 1, 2, 3}, 1))        //8
}
