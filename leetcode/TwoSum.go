package leetcode

func twoSum(nums []int, target int) []int {
	dictNum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dictNum[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if _, ok := dictNum[target-val]; ok && dictNum[target-val] != i {
			return []int{i, dictNum[target-val]}
		}
	}
	return nil
}
