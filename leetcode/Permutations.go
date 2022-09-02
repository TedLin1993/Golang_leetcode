package leetcode

import "fmt"

var permutations [][]int

func permute(nums []int) [][]int {
	permutations = [][]int{}
	backtracking_permute(nums, []int{})
	return permutations
}

func backtracking_permute(nums []int, permutation []int) {
	if len(nums) == 0 {
		p := append([]int{}, permutation...)
		permutations = append(permutations, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		nums = append(nums[:i], nums[i+1:]...)
		permutation = append(permutation, v)
		backtracking_permute(nums, permutation)
		permutation = permutation[:len(permutation)-1]
		nums = append(nums[:i+1], nums[i:]...)
		nums[i] = v
	}
}

func Test_permute() {
	fmt.Println(permute([]int{1, 2, 3}))    //[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
	fmt.Println(permute([]int{0, 1}))       //[[0,1],[1,0]]
	fmt.Println(permute([]int{1}))          //[[1]]
	fmt.Println(permute([]int{5, 4, 6, 2})) //

}
