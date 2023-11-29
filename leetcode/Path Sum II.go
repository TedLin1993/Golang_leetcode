package leetcode

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	var dfs func(node *TreeNode, path []int, sum int)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		path = append(path, node.Val)
		if sum == targetSum && node.Left == nil && node.Right == nil {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
		path = path[:len(path)-1]
	}
	dfs(root, []int{}, 0)
	return res
}
