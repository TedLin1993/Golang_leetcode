package leetcode

func rotate_189(nums []int, k int) {
	n := len(nums)
	k %= n
	arr := make([]int, n)
	copy(arr, nums)
	for i := 0; i < n; i++ {
		if i < k {
			nums[i] = arr[n-k+i]
		} else {
			nums[i] = arr[i-k]
		}
	}
}
