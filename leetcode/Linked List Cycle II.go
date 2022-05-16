package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head
	isCycle := false
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	//Floyd's algorithm: 起始node到迴圈node的距離，與slow/head相交node到迴圈node的距離等距
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
