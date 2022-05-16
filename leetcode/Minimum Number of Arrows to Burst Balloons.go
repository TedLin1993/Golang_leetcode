package leetcode

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 0
	interRange := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][0] < interRange[0] {
			continue
		} else if points[i][0] <= interRange[1] {
			interRange[0] = points[i][0]
		} else {
			res++
			interRange = points[i]
		}
	}
	res++
	return res
}

func TestMinimum_Number_of_Arrows_to_Burst_Balloons() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))                                                //2
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))                                                   //4
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))                                                   //2
	fmt.Println(findMinArrowShots([][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}})) //2

}
