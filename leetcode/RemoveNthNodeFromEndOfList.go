package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := ListNode{Val: 0, Next: head}
	left, right := &first, &first
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return first.Next
}
