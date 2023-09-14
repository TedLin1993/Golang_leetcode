package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for a NextTreeNode.
type NextTreeNode struct {
	Val   int
	Left  *NextTreeNode
	Right *NextTreeNode
	Next  *NextTreeNode
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sumArr(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Node struct {
	Val      int
	Children []*Node
}

func maxArr(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return res
}

func reverseByteArr(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
