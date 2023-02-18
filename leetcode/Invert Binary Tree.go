package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dfs_invertTree(root)
	return root
}

func dfs_invertTree(node *TreeNode) {
	node.Left, node.Right = node.Right, node.Left
	if node.Left != nil {
		dfs_invertTree(node.Left)
	}
	if node.Right != nil {
		dfs_invertTree(node.Right)
	}
}
