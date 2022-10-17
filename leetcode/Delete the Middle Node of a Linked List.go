package leetcode

import "fmt"

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func Test_deleteMiddle() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node1.Next = node2
	node2.Next = &node3
	node3.Next = &node4
	fmt.Println(deleteMiddle(node1))
}
