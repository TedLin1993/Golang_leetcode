package leetcode

func pathSum_III(root *TreeNode, targetSum int) int {
	preSum := map[int]int{0: 1}
	res := 0
	var dfs func(node *TreeNode, cur int)
	dfs = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		res += preSum[cur-targetSum]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]--
	}
	dfs(root, 0)
	return res
}
