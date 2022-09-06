package leetcode

import "fmt"

func pruneTree(root *TreeNode) *TreeNode {
	if !dfs_pruneTree(root) {
		return nil
	}
	return root
}

func dfs_pruneTree(node *TreeNode) bool {
	if node == nil {
		return false
	}

	if !dfs_pruneTree(node.Left) {
		node.Left = nil
	}
	if !dfs_pruneTree(node.Right) {
		node.Right = nil
	}
	return node.Left != nil || node.Right != nil || node.Val == 1
}

func Test_pruneTree() {
	fmt.Println(pruneTree(&TreeNode{Val: 0}))
}
