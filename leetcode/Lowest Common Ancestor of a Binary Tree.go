package leetcode

func lowestCommonAncestor_BT(currentNode, p, q *TreeNode) *TreeNode {
	if currentNode == nil {
		return nil
	}
	if currentNode == p || currentNode == q {
		return currentNode
	}

	left := lowestCommonAncestor(currentNode.Left, p, q)
	right := lowestCommonAncestor(currentNode.Right, p, q)

	//both left and right find target node, so currentNode is LCA
	if left != nil && right != nil {
		return currentNode
	}

	//if only one side find target node, it means one target node is under the other target node
	if left != nil {
		return left
	}
	return right
}
