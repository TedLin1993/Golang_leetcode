package leetcode

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		layer := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			layer = append(layer, queue[i].Val)
			for _, n := range queue[i].Children {
				queue = append(queue, n)
			}
		}
		res = append(res, layer)
		queue = queue[length:]
	}
	return res
}
