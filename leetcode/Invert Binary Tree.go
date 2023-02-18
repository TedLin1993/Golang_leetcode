package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := &TreeNode{Val: root.Val}
	node := res
	dfs_invertTree(root, node)
	return res
}

func dfs_invertTree(originNode, inverseNode *TreeNode) {
	if originNode.Left != nil {
		inverseNode.Right = &TreeNode{Val: originNode.Left.Val}
		dfs_invertTree(originNode.Left, inverseNode.Right)
	}
	if originNode.Right != nil {
		inverseNode.Left = &TreeNode{Val: originNode.Right.Val}
		dfs_invertTree(originNode.Right, inverseNode.Left)
	}
}
