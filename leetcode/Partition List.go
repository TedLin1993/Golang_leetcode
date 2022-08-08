package leetcode

import "fmt"

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy
	for right != nil && right.Next != nil {
		node := right.Next
		if node.Val >= x {
			right = right.Next
		} else if node.Val < x && left == right {
			left = left.Next
			right = right.Next
		} else {
			right.Next = right.Next.Next
			node.Next = left.Next
			left.Next = node
			left = left.Next
		}
	}
	return dummy.Next
}
func Testpartition() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 0}
	node5 := &ListNode{Val: 2}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	fmt.Println(partition(node1, 3))
}
