package leetcode

var sum int

func maxPathSum(root *TreeNode) int {
	sum = -1001
	dfsMaxPathSum(root)
	return sum
}

func dfsMaxPathSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSum := max(dfsMaxPathSum(node.Left), 0)
	rightSum := max(dfsMaxPathSum(node.Right), 0)

	sum = max(sum, leftSum+rightSum+node.Val)

	return node.Val + max(leftSum, rightSum)
}
