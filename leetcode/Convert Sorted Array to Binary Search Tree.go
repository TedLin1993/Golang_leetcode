package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}
	root.Left = sortedArrayToBST(nums[:midIndex])
	root.Right = sortedArrayToBST(nums[midIndex+1:])
	return root
}

func Test_sortedArrayToBST() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
