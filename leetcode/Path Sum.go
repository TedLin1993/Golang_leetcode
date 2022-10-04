package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs_hasPathSum(root, targetSum)
}

func dfs_hasPathSum(node *TreeNode, sum int) bool {
	if node.Left == nil && node.Right == nil {
		return sum-node.Val == 0
	}

	sum -= node.Val
	if node.Left != nil && dfs_hasPathSum(node.Left, sum) {
		return true
	}
	if node.Right != nil && dfs_hasPathSum(node.Right, sum) {
		return true
	}
	return false
}
