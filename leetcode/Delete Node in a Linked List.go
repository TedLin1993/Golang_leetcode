package leetcode

func deleteNode(node *ListNode) {
	*node = *node.Next
}
