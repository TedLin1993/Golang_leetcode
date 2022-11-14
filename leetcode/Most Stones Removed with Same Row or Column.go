package leetcode

import "fmt"

var islandMap map[int]int
var islandCount int

func removeStones(stones [][]int) int {
	islandMap = make(map[int]int, len(stones))
	islandCount = 0
	columnAdd := 1001
	for i := 0; i < len(stones); i++ {
		union(stones[i][0], stones[i][1]+columnAdd)
	}
	return len(stones) - islandCount
}

func union(a, b int) {
	a = find(a)
	b = find(b)
	if a != b {
		islandMap[a] = b
		islandCount--
	}
}

func find(a int) int {
	if islandMap[a] == 0 {
		islandMap[a] = a
		islandCount++
	}
	if a != islandMap[a] {
		islandMap[a] = find(islandMap[a])
	}
	return islandMap[a]
}

func Test_removeStones() {
	// fmt.Println(removeStones([][]int{{0, 0}}))                                                                                                                                                         //0
	fmt.Println(removeStones([][]int{{0, 2}, {0, 2}, {1, 1}, {2, 0}, {2, 2}})) //3
	// fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))                                                                                                                 //5
	// fmt.Println(removeStones([][]int{{5, 9}, {9, 0}, {0, 0}, {7, 0}, {4, 3}, {8, 5}, {5, 8}, {1, 1}, {0, 6}, {7, 5}, {1, 6}, {1, 9}, {9, 4}, {2, 8}, {1, 3}, {4, 2}, {2, 5}, {4, 1}, {0, 2}, {6, 5}})) //19
}
