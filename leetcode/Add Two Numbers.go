package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	carryValue := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		value := n1 + n2 + carryValue

		carryValue = value / 10
		value = value % 10
		node.Next = &ListNode{Val: value}
		node = node.Next
	}
	if carryValue == 1 {
		node.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
