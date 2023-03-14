package leetcode

func sumNumbers(root *TreeNode) int {
	return dfs_sumNumbers(root, 0)
}

func dfs_sumNumbers(node *TreeNode, num int) int {
	num = num*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return num
	}

	res := 0
	if node.Left != nil {
		res += dfs_sumNumbers(node.Left, num)
	}
	if node.Right != nil {
		res += dfs_sumNumbers(node.Right, num)
	}
	return res
}
