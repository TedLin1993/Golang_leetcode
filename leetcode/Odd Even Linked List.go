package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	left, right := head, head.Next
	for right != nil && right.Next != nil {
		nextOdd := right.Next
		right.Next = nextOdd.Next
		nextOdd.Next = left.Next
		left.Next = nextOdd
		left = left.Next
		right = right.Next
	}
	return head
}
