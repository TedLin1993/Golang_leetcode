package leetcode

func insertionSortList(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	node := head
	for node.Next != nil {
		if node.Next.Val < node.Val {
			formerNode := dummyNode
			temp := node.Next
			for {
				if formerNode.Next.Val > temp.Val {
					node.Next = temp.Next
					temp.Next = formerNode.Next
					formerNode.Next = temp
					break
				}
				formerNode = formerNode.Next
			}
		} else {
			node = node.Next
		}
	}
	return dummyNode.Next
}

func TestInsertionSortList() {
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 1}
	node4 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	insertionSortList(node1)
}
