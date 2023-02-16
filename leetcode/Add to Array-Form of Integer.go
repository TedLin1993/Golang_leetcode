package leetcode

import "fmt"

func addToArrayForm(nums []int, k int) []int {
	kLength, temp_k := 0, k
	for temp_k > 0 {
		temp_k = temp_k / 10
		kLength++
	}
	length := max(len(nums), kLength) + 1
	res := make([]int, length)
	carry := 0
	idx, numsIdx := length-1, len(nums)-1
	for idx >= 0 {
		sum := carry
		if numsIdx >= 0 {
			sum += nums[numsIdx]
			numsIdx--
		}
		if k > 0 {
			sum += k % 10
			k /= 10
		}
		carry = sum / 10
		res[idx] = sum % 10
		idx--
	}
	if res[0] == 0 {
		res = res[1:]
	}
	return res
}

func Test_addToArrayForm() {
	fmt.Println(addToArrayForm([]int{1, 2, 3, 4}, 123)) //[1 3 5 7]
	fmt.Println(addToArrayForm([]int{9, 8, 7}, 321))    //[1 3 0 8]
}
