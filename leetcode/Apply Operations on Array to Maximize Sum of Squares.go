package leetcode

func maxSum(nums []int, k int) int {
	bitCount := make([]int, 32)
	for _, num := range nums {
		i := 0
		for num > 0 {
			if num%2 == 1 {
				bitCount[i]++
			}
			num = num >> 1
			i++
		}
	}

	res := 0
	mod := int(1e9 + 7)
	for i := 0; i < k; i++ {
		temp := 0
		for j := 0; j < 32; j++ {
			if bitCount[j] > 0 {
				temp += 1 << j
				bitCount[j]--
			}
		}
		temp = temp * temp % mod
		res = (res + temp) % mod
	}
	return res
}
