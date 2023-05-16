package leetcode

var nodeSet []int

func countCompleteComponents(n int, edges [][]int) int {
	nodeSet = make([]int, n)
	edgeCountMap := make([]int, n)
	for i := 0; i < n; i++ {
		nodeSet[i] = i
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		union_countCompleteComponents(a, b)
		edgeCountMap[a]++
		edgeCountMap[b]++
	}
	for i := 0; i < n; i++ {
		find_countCompleteComponents(i)
	}

	res := 0
	for i := 0; i < n; i++ {
		if nodeSet[i] != i {
			continue
		}

		isComplete := true
		count := getCount(nodeSet, i) - 1
		for j := i; j < n; j++ {
			if nodeSet[j] == i {
				if edgeCountMap[j] != count {
					isComplete = false
					break
				}
			}
		}
		if isComplete {
			res++
		}
	}

	return res
}

func find_countCompleteComponents(idx int) int {
	if idx == nodeSet[idx] {
		return idx
	}
	nodeSet[idx] = find_countCompleteComponents(nodeSet[idx])
	return nodeSet[idx]
}

func union_countCompleteComponents(a, b int) {
	a, b = find_countCompleteComponents(a), find_countCompleteComponents(b)
	if a < b {
		nodeSet[b] = a
	} else {
		nodeSet[a] = b
	}
}

func getCount(arr []int, target int) int {
	res := 0
	for _, a := range arr {
		if a == target {
			res++
		}
	}
	return res
}
