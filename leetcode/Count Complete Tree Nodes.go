package leetcode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	node := root
	for node != nil {
		leftDepth := getDepth(node.Left)
		rightDepth := getDepth(node.Right)
		if leftDepth == rightDepth {
			res += 1 << leftDepth
			node = node.Right
		} else {

		}
	}
	return res
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + getDepth(node.Left)
}
