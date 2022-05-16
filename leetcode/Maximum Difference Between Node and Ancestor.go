package leetcode

func maxAncestorDiff(root *TreeNode) int {
	max, min := root.Val, root.Val
	return dfsMaxAncestorDiff(max, min, root)
}

func dfsMaxAncestorDiff(max, min int, node *TreeNode) int {
	if node == nil {
		return max - min
	}

	if node.Val > max {
		max = node.Val
	} else if node.Val < min {
		min = node.Val
	}

	leftVal := dfsMaxAncestorDiff(max, min, node.Left)
	rightVal := dfsMaxAncestorDiff(max, min, node.Right)
	if leftVal > rightVal {
		return leftVal
	} else {
		return rightVal
	}
}
