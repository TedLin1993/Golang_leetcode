package leetcode

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	flightMap := make([][]int, n)
	for i := 0; i < n; i++ {
		flightMap[i] = make([]int, n)
		for j := 0; j < n; j++ {
			flightMap[i][j] = math.MaxInt
		}
	}
	for _, flight := range flights {
		from := flight[0]
		to := flight[1]
		cost := flight[2]
		flightMap[from][to] = cost
	}
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt
	}
	queue := []int{src}
	cost[src] = 0
	for len(queue) > 0 {
		if k < 0 {
			break
		}
		length := len(queue)
		tempCost := make([]int, n)
		for i := 0; i < n; i++ {
			tempCost[i] = math.MaxInt
		}
		for i := 0; i < length; i++ {
			node := queue[i]
			for i := 0; i < n; i++ {
				if flightMap[node][i] == math.MaxInt {
					continue
				}
				temp := cost[node] + flightMap[node][i]
				if temp < cost[i] && temp < tempCost[i] {
					tempCost[i] = temp
					queue = append(queue, i)
				}
			}
		}
		for i := 0; i < n; i++ {
			if tempCost[i] < math.MaxInt {
				cost[i] = tempCost[i]
			}
		}
		queue = queue[length:]
		k--
	}

	if cost[dst] != math.MaxInt {
		return cost[dst]
	}
	return -1
}
