package leetcode

import "math"

var res_minDiffInBST int

func minDiffInBST(root *TreeNode) int {
	res_minDiffInBST = math.MaxInt32
	dfs_minDiffInBST(root.Left, math.MinInt32, root.Val)
	dfs_minDiffInBST(root.Right, root.Val, math.MaxInt32)
	return res_minDiffInBST
}

func dfs_minDiffInBST(node *TreeNode, left, right int) {
	if node == nil {
		return
	}
	res_minDiffInBST = min(res_minDiffInBST, node.Val-left)
	res_minDiffInBST = min(res_minDiffInBST, right-node.Val)
	dfs_minDiffInBST(node.Left, left, node.Val)
	dfs_minDiffInBST(node.Right, node.Val, right)
}
