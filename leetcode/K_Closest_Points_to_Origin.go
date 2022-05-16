package leetcode

import (
	"fmt"
)

func kClosest(points [][]int, k int) [][]int {
	distances := []int{}
	min, max := 0, 0
	pointsPos := []int{}
	for i := 0; i < len(points); i++ {
		dist := getDistance(points[i])
		distances = append(distances, dist)
		if dist > max {
			max = dist
		}
		pointsPos = append(pointsPos, i)
	}
	resultPos := []int{}
	for k > 0 {
		mid := min + (max-min)/2
		leftPointsPos, rightPointsPos := splitPoints(pointsPos, mid, distances)
		if len(leftPointsPos) > k {
			pointsPos = leftPointsPos
			max = mid
		} else {
			k -= len(leftPointsPos)
			pointsPos = rightPointsPos
			min = mid
			resultPos = append(resultPos, leftPointsPos...)
		}
	}

	result := [][]int{}
	for i := 0; i < len(resultPos); i++ {
		result = append(result, points[resultPos[i]])
	}
	return result
}

func getDistance(point []int) int {
	x := point[0]
	y := point[1]
	return x*x + y*y
}

func splitPoints(pointPos []int, mid int, distances []int) (leftPointsPos, rightPointsPos []int) {
	for i := 0; i < len(pointPos); i++ {
		if distances[pointPos[i]] < mid {
			leftPointsPos = append(leftPointsPos, pointPos[i])
		} else {
			rightPointsPos = append(rightPointsPos, pointPos[i])
		}
	}
	return
}

func TestK_Closest_Points_to_Origin() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))          //[[-2,2]]
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)) //[[3,3],[-2,4]]
}
