package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	isFromLeft := true
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		count := len(queue)
		temp := make([]int, count)
		for i := 0; i < count; i++ {
			if isFromLeft {
				temp[i] = queue[i].Val
			} else {
				temp[count-1-i] = queue[i].Val
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, temp)
		isFromLeft = !isFromLeft
		queue = queue[count:]
	}
	return res
}
