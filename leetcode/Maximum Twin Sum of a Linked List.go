package leetcode

func pairSum(head *ListNode) int {
	dummy := &ListNode{Next: head}
	mid, node := dummy, dummy
	for node.Next != nil && node.Next.Next != nil {
		mid = mid.Next
		node = node.Next.Next
	}
	nextHead := mid.Next
	//reverse previous half nodes
	left, right := dummy, dummy.Next
	for right != nextHead {
		temp := right.Next
		right.Next = left
		left, right = right, temp
	}

	res := 0
	node1, node2 := mid, nextHead
	for node1 != nil && node2 != nil {
		res = max(res, node1.Val+node2.Val)
		node1 = node1.Next
		node2 = node2.Next
	}
	return res
}
