package leetcode

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	node := dummyNode
	for node.Next != nil && node.Next.Next != nil {
		temp := node.Next.Next
		node.Next.Next = temp.Next
		temp.Next = node.Next
		node.Next = temp
		node = node.Next.Next
	}

	return dummyNode.Next
}
