package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	if !dfsIsValidBST(root.Left, math.MinInt64, root.Val) {
		return false
	}
	if !dfsIsValidBST(root.Right, root.Val, math.MaxInt64) {
		return false
	}
	return true
}

func dfsIsValidBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}

	if !dfsIsValidBST(node.Left, min, node.Val) {
		return false
	}
	if !dfsIsValidBST(node.Right, node.Val, max) {
		return false
	}
	return true
}
