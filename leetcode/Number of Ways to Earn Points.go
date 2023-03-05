package leetcode

import (
	"fmt"
	"sort"
)

func waysToReachTarget(target int, types [][]int) int {
	sum := 0
	for _, t := range types {
		sum += t[0] * t[1]
	}
	if sum < target {
		return 0
	}
	if sum == target {
		return 1
	}
	sort.Slice(types, func(i, j int) bool {
		return types[i][1] > types[j][1]
	})
	return dfs_waysToReachTarget(target,
		
		
		
		
		
		
		types, 0)
}

func dfs_waysToReachTarget(target int, types [][]int, typeIdx int) int {
	if typeIdx >= len(types) {
		return 0
	}
	modulo := 1000000007
	typeCount := types[typeIdx][0]
	typeMarks := types[typeIdx][1]
	res := dfs_waysToReachTarget(target, types, typeIdx+1)
	for i := 0; i < typeCount; i++ {
		target -= typeMarks
		if target == 0 {
			return res + 1
		}
		if target < 0 {
			return res
		}
		res += dfs_waysToReachTarget(target, types, typeIdx+1)
		res %= modulo
	}
	return res
}

func Test_waysToReachTarget() {
	fmt.Println(waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))    //7
	fmt.Println(waysToReachTarget(5, [][]int{{50, 1}, {50, 2}, {50, 5}})) //4
	fmt.Println(waysToReachTarget(18, [][]int{{6, 1}, {3, 2}, {2, 3}}))   //1
}
