package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil
	head = sortList(head)
	head2 = sortList(head2)
	return mergeList(head, head2)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return dummy.Next
}
