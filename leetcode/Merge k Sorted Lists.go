package leetcode

import "container/heap"

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

func mergeKLists3(lists []*ListNode) *ListNode {
	listHeap := NodeHeap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&listHeap, lists[i])
		}
	}
	dummy := &ListNode{}
	currentNode := dummy
	for len(listHeap) != 0 {
		node := heap.Pop(&listHeap).(*ListNode)
		currentNode.Next = node
		currentNode = currentNode.Next
		if node.Next != nil {
			heap.Push(&listHeap, node.Next)
		}
	}
	return dummy.Next
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
