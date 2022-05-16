package leetcode

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	gasTotal := 0
	costTotal := 0
	gasRemand := 0
	start := 0
	for i := 0; i < length; i++ {
		gasTotal += gas[i]
		costTotal += cost[i]

		gasRemand = gasRemand + gas[i] - cost[i]
		if gasRemand < 0 {
			start = i + 1
			gasRemand = 0
		}
	}
	if gasTotal < costTotal {
		return -1
	}
	return start
}

func TestCanCompleteCircuit() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) //3
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             //-1
}
