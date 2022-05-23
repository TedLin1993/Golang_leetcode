package leetcode

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 1
	lastNode := head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		length++
	}
	
	k = k % length
	if k == 0 {
		return head
	}

	newLastNode := head
	for i := 0; i < length-k-1; i++ {
		newLastNode = newLastNode.Next
	}

	newHead := newLastNode.Next
	newLastNode.Next = nil
	lastNode.Next = head
	return newHead
}

func TestrotateRight() {
	node1 := &ListNode{Val: 0}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node3

	fmt.Println(rotateRight(node1, 3))
}
