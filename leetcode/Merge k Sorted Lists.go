package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var res *ListNode
	for i := 0; i < len(lists); i++ {
		res = merge2List(res, lists[i])
	}

	return res
}

func merge2List(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val < right.Val {
		left.Next = merge2List(left.Next, right)
		return left
	}
	right.Next = merge2List(left, right.Next)
	return right
}
