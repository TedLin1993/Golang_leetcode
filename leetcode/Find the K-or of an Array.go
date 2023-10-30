package leetcode

func findKOr(nums []int, k int) int {
	bitCount := make([]int, 32)
	for _, num := range nums {
		i := 0
		for num > 0 {
			if num%2 == 1 {
				bitCount[i]++
			}
			i++
			num = num >> 1
		}
	}
	res := 0
	for i := 0; i < 32; i++ {
		if bitCount[i] >= k {
			res += 1 << i
		}
	}
	return res
}
