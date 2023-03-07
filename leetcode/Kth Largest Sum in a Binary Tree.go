package leetcode

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levelSums := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		sum := 0
		for i := 0; i < length; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		levelSums = append(levelSums, sum)
		queue = queue[length:]
	}
	if k > len(levelSums) {
		return -1
	}
	res := quickSelect(levelSums, k)
	return int64(res)
}

func quickSelect(nums []int, k int) int {
	return quickSelectKthLargest(nums, len(nums)-k, 0, len(nums)-1)
}

func quickSelectKthLargest(nums []int, targetIdx int, left, right int) int {
	pivot := nums[right]
	pivotIdx := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[pivotIdx], nums[i] = nums[i], nums[pivotIdx]
			pivotIdx++
		}
	}
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	if pivotIdx == targetIdx {
		return pivot
	}
	if pivotIdx < targetIdx {
		return quickSelectKthLargest(nums, targetIdx, pivotIdx+1, right)
	} else {
		return quickSelectKthLargest(nums, targetIdx, left, pivotIdx-1)
	}
}
