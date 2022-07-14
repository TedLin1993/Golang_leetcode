package leetcode

import "fmt"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		len := len(queue)
		res = append(res, queue[len-1].Val)
		for i := 0; i < len; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[len:]
	}
	return res
}

func TestrightSideView() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node3.Right = &node5
	fmt.Println(rightSideView(&node1))
}
