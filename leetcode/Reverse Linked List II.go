package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	preLeft := &ListNode{Next: head}
	node := preLeft
	for i := 0; i < left-1; i++ {
		node = node.Next
	}
	preLeft = node
	node = node.Next

	for i := 0; i < right-left; i++ {
		next := node.Next
		node.Next = next.Next
		next.Next = preLeft.Next
		preLeft.Next = next
	}
	if left == 1 {
		return preLeft.Next
	}
	return head
}
