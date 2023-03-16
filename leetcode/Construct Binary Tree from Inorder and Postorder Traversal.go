package leetcode

func buildTree_InorderPostorder(inorder []int, postorder []int) *TreeNode {
	valMap := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		valMap[val] = idx
	}
	var dfs func(inStart, inEnd, postStart, postEnd int) *TreeNode
	dfs = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		headVal := postorder[postEnd]
		headIdx := valMap[headVal]
		delta := headIdx - inStart
		head := &TreeNode{Val: headVal}
		head.Left = dfs(inStart, headIdx-1, postStart, postStart+delta-1)
		head.Right = dfs(headIdx+1, inEnd, postStart+delta, postEnd-1)
		return head
	}
	return dfs(0, len(inorder)-1, 0, len(postorder)-1)
}
