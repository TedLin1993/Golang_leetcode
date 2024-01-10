package leetcode

func amountOfTime(root *TreeNode, start int) int {
	graph := map[int][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	visited := map[int]bool{}
	queue2 := []int{start}
	res := -1
	for len(queue2) > 0 {
		count := len(queue2)
		for i := 0; i < count; i++ {
			cur := queue2[i]
			visited[cur] = true
			for _, next := range graph[cur] {
				if visited[next] {
					continue
				}
				queue2 = append(queue2, next)
			}
		}
		queue2 = queue2[count:]
		res++
	}
	return res
}
