package leetcode

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	res := root.Val
	for len(queue) > 0 {
		res = queue[0].Val
		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[count:]
	}
	return res
}
