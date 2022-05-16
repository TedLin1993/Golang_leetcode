package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	for {
		if val < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return root
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				return root
			}
			node = node.Right
		}
	}
}

func TestInsert_into_a_Binary_Search_Tree() {

}
