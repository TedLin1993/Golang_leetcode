package leetcode

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		sum := 0.0
		for i := 0; i < length; i++ {
			sum += float64(queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
		res = append(res, sum/float64(length))
	}
	return res
}
