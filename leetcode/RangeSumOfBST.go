package leetcode

func rangeSumBST(root *TreeNode, low int, high int) int {
	return dfsSumBST(root, low, high)
}

func dfsSumBST(node *TreeNode, low, high int) int {
	if node == nil {
		return 0
	}
	if node.Val < low {
		return dfsSumBST(node.Right, low, high)
	} else if node.Val > high {
		return dfsSumBST(node.Left, low, high)
	} else {
		return node.Val + dfsSumBST(node.Right, low, high) + dfsSumBST(node.Left, low, high)
	}
}
