package leetcode

import (
	"fmt"
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	dp := 1
	factorMap := map[int]int{arr[0]: 1}
	mod := 1000000007
	for i := 1; i < len(arr); i++ {
		parent := arr[i]
		for j := 0; j < i; j++ {
			child := arr[j]
			if parent%child != 0 {
				continue
			}
			//another child can divide parent, and it should not less than child
			v := parent / child
			if v < child {
				break
			} else if v == child {
				//if the two children are the same, then they change position are not diffrent
				factorMap[parent] += factorMap[child] * factorMap[v]
			} else {
				//we can change the two children's position, so we should multiply by 2
				factorMap[parent] += factorMap[child] * factorMap[v] * 2
			}
		}
		//only itself can be seen as a tree
		factorMap[parent] += 1
		dp += factorMap[parent]
		dp %= mod
	}
	return dp
}

func Test_numFactoredBinaryTrees() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 4}))            //3
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))     //7
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10, 20})) //18
}
