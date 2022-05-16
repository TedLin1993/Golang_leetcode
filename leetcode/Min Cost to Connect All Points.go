package leetcode

import (
	"fmt"
	"math"
)

func minCostConnectPoints(points [][]int) int {
	if len(points) == 1 {
		return 0
	}
	res := 0

	minDistNodes := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		minDistNodes[i] = math.MaxInt32
	}
	currentNode := 0
	nodeVisited := make([]bool, len(points))
	nodeVisited[0] = true
	for i := 1; i < len(points); i++ {
		minNode := -1
		minDist := math.MaxInt32
		for node := 0; node < len(points); node++ {
			if nodeVisited[node] {
				continue
			}
			dist := abs(points[currentNode][0]-points[node][0]) +
				abs(points[currentNode][1]-points[node][1])
			if dist < minDistNodes[node] {
				minDistNodes[node] = dist
			}

			if minDistNodes[node] < minDist {
				minNode = node
				minDist = minDistNodes[node]
			}
		}
		res += minDist
		currentNode = minNode
		nodeVisited[currentNode] = true
	}
	return res
}

func TestminCostConnectPoints() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))  //20
	fmt.Println(minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))                //18
	fmt.Println(minCostConnectPoints([][]int{{-1000000, -1000000}, {1000000, 1000000}})) //4000000
}
