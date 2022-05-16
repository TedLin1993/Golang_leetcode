package leetcode

import "fmt"

func increasingBST(root *TreeNode) *TreeNode {
	return inorderBST(root, nil)
}

func inorderBST(root, nextNode *TreeNode) *TreeNode {
	if root == nil {
		return nextNode
	}

	res := inorderBST(root.Left, root)
	root.Left = nil
	root.Right = inorderBST(root.Right, nextNode)

	return res
}

func TestincreasingBST() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node3.Right = &node5
	fmt.Println(increasingBST(&node1))
}
