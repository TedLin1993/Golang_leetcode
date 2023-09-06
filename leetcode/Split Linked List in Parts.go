package leetcode

import "fmt"

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := 0
	node := head
	for node != nil {
		count++
		node = node.Next
	}

	avg := count / k
	curK := k
	node = head
	res := make([]*ListNode, 0, k)
	for count > 0 {
		res = append(res, node)
		length := avg
		if count%curK > 0 {
			length++
		}
		temp := node
		for i := 0; i < length-1; i++ {
			temp = temp.Next
		}
		node = temp.Next
		temp.Next = nil
		curK--
		count -= length
	}
	for len(res) < k {
		res = append(res, nil)
	}
	return res
}

func Test_splitListToParts() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	fmt.Println(splitListToParts(node1, 3))
}
