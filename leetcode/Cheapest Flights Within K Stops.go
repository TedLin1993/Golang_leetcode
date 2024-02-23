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

func findCheapestPrice_2(n int, flights [][]int, src int, dst int, k int) int {
	type fly struct {
		to    int
		price int
	}
	graph := make(map[int][]fly)
	for _, f := range flights {
		graph[f[0]] = append(graph[f[0]], fly{f[1], f[2]})
	}
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt
	}
	costs[src] = 0
	queue := []int{src}
	for len(queue) > 0 && k >= 0 {
		temp := make([]int, n)
		copy(temp, costs)
		count := len(queue)
		for i := 0; i < count; i++ {
			city := queue[i]
			for _, f := range graph[city] {
				if costs[city]+f.price < temp[f.to] {
					temp[f.to] = costs[city] + f.price
					queue = append(queue, f.to)
				}
			}
		}
		costs = temp
		queue = queue[count:]
		k--
	}
	if costs[dst] == math.MaxInt {
		return -1
	}
	return costs[dst]
}
