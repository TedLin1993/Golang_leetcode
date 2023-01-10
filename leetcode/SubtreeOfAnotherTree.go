package leetcode

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if isSubSameTree(node, subRoot) {
			return true
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}

func isSubSameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}

	return isSubSameTree(a.Left, b.Left) && isSubSameTree(a.Right, b.Right)
}
