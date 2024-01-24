package leetcode

func pseudoPalindromicPaths(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, valCount int)
	dfs = func(node *TreeNode, valCount int) {
		valCount = valCount ^ (1 << node.Val)
		if node.Left == nil && node.Right == nil {
			if valCount&(valCount-1) == 0 {
				res++
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, valCount)
		}
		if node.Right != nil {
			dfs(node.Right, valCount)
		}
	}
	dfs(root, 0)
	return res
}
