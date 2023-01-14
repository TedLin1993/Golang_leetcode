package leetcode

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	res := 0
	for i := 0; i < len(points)-1; i++ {
		vectorMap := make(map[float32]int, len(points))
		for j := i + 1; j < len(points); j++ {
			v_x := points[j][0] - points[i][0]
			v_y := points[j][1] - points[i][1]
			var vector float32
			if v_x == 0 {
				vector = math.MaxInt
			} else {
				vector = float32(v_y) / float32(v_x)
			}
			vectorMap[vector]++
		}
		for _, v := range vectorMap {
			if v+1 > res {
				res = v + 1
			}
		}
	}
	return res
}

func Test_maxPoints() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))                                 //3
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))         //4
	fmt.Println(maxPoints([][]int{{0, 0}, {1, -1}, {1, 1}}))                                //2
	fmt.Println(maxPoints([][]int{{0, 0}, {4, 5}, {7, 8}, {8, 9}, {5, 6}, {3, 4}, {1, 1}})) //5
}
