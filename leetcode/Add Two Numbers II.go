package leetcode

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverseList(l1), reverseList(l2)
	head := &ListNode{}
	node := head
	carry := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		node.Val = val % 10
		carry = val / 10
		l1, l2 = l1.Next, l2.Next
		node.Next = &ListNode{}
		node = node.Next
	}
	for l1 != nil {
		val := l1.Val + carry
		node.Val = val % 10
		carry = val / 10
		l1 = l1.Next
		node.Next = &ListNode{}
		node = node.Next
	}
	for l2 != nil {
		val := l2.Val + carry
		node.Val = val % 10
		carry = val / 10
		l2 = l2.Next
		node.Next = &ListNode{}
		node = node.Next
	}
	if carry == 1 {
		node.Val = 1
	}
	head = reverseList(head)
	if head.Val == 0 {
		head = head.Next
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	left, right := head, head.Next
	left.Next = nil
	for right != nil {
		temp := right.Next
		right.Next = left
		left, right = right, temp
	}
	return left
}
