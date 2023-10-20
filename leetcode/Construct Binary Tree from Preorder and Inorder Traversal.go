package leetcode

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}

	cur := TreeNode{Val: preorder[0]}
	curIdx := 0
	for i := 0; i < n; i++ {
		if inorder[i] == cur.Val {
			curIdx = i
			break
		}
	}
	cur.Left = buildTree(preorder[1:curIdx+1], inorder[:curIdx])
	cur.Right = buildTree(preorder[curIdx+1:], inorder[curIdx+1:])
	return &cur
}

func TestbuildTree() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
