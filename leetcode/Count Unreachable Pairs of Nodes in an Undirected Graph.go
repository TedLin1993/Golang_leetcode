package leetcode

var originMap []int

func countPairs(n int, edges [][]int) int64 {
	originMap = make([]int, n)
	for idx := range originMap {
		originMap[idx] = idx
	}
	for _, edge := range edges {
		union(edge[0], edge[1])
	}
	for i := 0; i < n; i++ {
		find(i)
	}

	countMap := make(map[int]int64, n)
	var totalCount int64
	for _, origin := range originMap {
		countMap[origin]++
		totalCount++
	}

	var res int64
	for _, count := range countMap {
		totalCount -= count
		res += count * totalCount
	}
	return res
}

func countPairs_find(node int) int {
	if originMap[node] == node {
		return node
	}
	originMap[node] = find(originMap[node])
	return originMap[node]
}

func countPairs_union(node1, node2 int) {
	node1, node2 = find(node1), find(node2)
	if node1 < node2 {
		originMap[node2] = node1
	} else {
		originMap[node1] = node2
	}
}
