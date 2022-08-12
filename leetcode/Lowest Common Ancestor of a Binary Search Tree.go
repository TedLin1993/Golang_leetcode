package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var left, right *TreeNode
	if p.Val < q.Val {
		left = p
		right = q
	} else {
		left = q
		right = p
	}
	return dfsLowestCommonAncestor(root, left, right)
}

func dfsLowestCommonAncestor(node, left, right *TreeNode) *TreeNode {
	if (node.Val > left.Val && node.Val < right.Val) || node == left || node == right {
		return node
	}
	if node.Val > right.Val {
		return dfsLowestCommonAncestor(node.Left, left, right)
	}
	return dfsLowestCommonAncestor(node.Right, left, right)
}
