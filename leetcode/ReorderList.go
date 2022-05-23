package leetcode

func reorderList(head *ListNode) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}
	//Find the middle of the list
	middleNode := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		middleNode = middleNode.Next
	}

	//Reverse the half after middle 1->2->3->4->5->6 to 1->2->3->6->5->4
	node := middleNode.Next
	for node.Next != nil {
		temp := node.Next
		node.Next = temp.Next
		temp.Next = middleNode.Next
		middleNode.Next = temp
	}

	//Start reorder one by one 1->2->3->6->5->4 to 1->6->2->5->3->4
	leftNode := head
	for leftNode != middleNode {
		temp := middleNode.Next
		middleNode.Next = temp.Next
		temp.Next = leftNode.Next
		leftNode.Next = temp
		leftNode = temp.Next
	}
}

func TestReorderList() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	// node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	// node4.Next = node5
	reorderList(node1)
}
