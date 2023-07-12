package leetcode

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	max := 500
	nodeMap := make([][]int, max+1)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			nodeMap[node.Val] = append(nodeMap[node.Val], node.Left.Val)
			nodeMap[node.Left.Val] = append(nodeMap[node.Left.Val], node.Val)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			nodeMap[node.Val] = append(nodeMap[node.Val], node.Right.Val)
			nodeMap[node.Right.Val] = append(nodeMap[node.Right.Val], node.Val)
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}

	queueVal := []int{target.Val}
	visited := make([]bool, max+1)
	for k > 0 {
		length := len(queueVal)
		for i := 0; i < length; i++ {
			visited[queueVal[i]] = true
			for _, val := range nodeMap[queueVal[i]] {
				if !visited[val] {
					queueVal = append(queueVal, val)
				}
			}
		}
		queueVal = queueVal[length:]
		k--
	}
	return queueVal
}
