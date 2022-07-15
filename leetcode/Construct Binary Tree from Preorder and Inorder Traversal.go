package leetcode

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	node := &TreeNode{Val: preorder[0]}
	rootInorderIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootInorderIndex = i
			break
		}
	}
	if rootInorderIndex >= 1 {
		node.Left = buildTree(preorder[1:rootInorderIndex+1], inorder[:rootInorderIndex])
	}
	if rootInorderIndex < len(preorder)-1 {
		node.Right = buildTree(preorder[rootInorderIndex+1:], inorder[rootInorderIndex+1:])
	}

	return node
}

func TestbuildTree() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
