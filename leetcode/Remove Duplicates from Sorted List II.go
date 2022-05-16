package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	node := dummy
	for node != nil && node.Next != nil {
		nextNode := node.Next
		if nextNode.Next != nil && nextNode.Val == nextNode.Next.Val {
			for nextNode.Next != nil && nextNode.Val == nextNode.Next.Val {
				nextNode = nextNode.Next
			}
			nextNode = nextNode.Next
			node.Next = nextNode
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}
