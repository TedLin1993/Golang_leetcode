package leetcode

import "fmt"

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	nodeMap := make([]int, n)
	for i := range nodeMap {
		nodeMap[i] = i
	}
	var find func(node int) int
	find = func(node int) int {
		if nodeMap[node] == node {
			return node
		}
		nodeMap[node] = find(nodeMap[node])
		return nodeMap[node]
	}
	union := func(node1, node2 int) {
		node1, node2 = find(node1), find(node2)
		nodeMap[node2] = node1
	}

	for i, bomb1 := range bombs[:n-1] {
		for j := i + 1; j < n; j++ {
			bomb2 := bombs[j]
			disSqure := getDistanceSqare(bomb1, bomb2)
			if bomb1[2]*bomb1[2] >= disSqure {
				union(i, j)
			} else if bomb2[2]*bomb2[2] >= disSqure {
				union(j, i)
			}
		}
	}

	for i := range nodeMap {
		find(i)
	}
	return findMaxFrequent(nodeMap)
}

func getDistanceSqare(a, b []int) int {
	dist1, dist2 := a[0]-b[0], a[1]-b[1]
	return dist1*dist1 + dist2*dist2
}

func findMaxFrequent(array []int) int {
	freqMap := make(map[int]int)

	for _, element := range array {
		freqMap[element]++
	}

	maxFreq := 0

	for _, frequency := range freqMap {
		if frequency > maxFreq {
			maxFreq = frequency
		}
	}

	return maxFreq
}

func Test_maximumDetonation() {
	// fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
	fmt.Println(maximumDetonation([][]int{{855, 82, 158}, {17, 719, 430}, {90, 756, 164}, {376, 17, 340}, {691, 636, 152}, {565, 776, 5}, {464, 154, 271}, {53, 361, 162}, {278, 609, 82}, {202, 927, 219}, {542, 865, 377}, {330, 402, 270}, {720, 199, 10}, {986, 697, 443}, {471, 296, 69}, {393, 81, 404}, {127, 405, 177}}))
}
