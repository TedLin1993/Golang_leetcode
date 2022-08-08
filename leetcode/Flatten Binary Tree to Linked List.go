package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dfsFlatten(root)
}

// flatten nodes and return last node
func dfsFlatten(currentNode *TreeNode) (lastNode *TreeNode) {
	if currentNode == nil {
		return nil
	}
	if currentNode.Left == nil && currentNode.Right == nil {
		return currentNode
	}

	leftLast := dfsFlatten(currentNode.Left)
	rightLast := dfsFlatten(currentNode.Right)
	if leftLast == nil {
		return rightLast
	}
	if rightLast == nil {
		currentNode.Right = currentNode.Left
		currentNode.Left = nil
		return leftLast
	}
	leftLast.Right = currentNode.Right
	currentNode.Right = currentNode.Left
	currentNode.Left = nil
	return rightLast
}
