package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}

		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	return nodeA
}
