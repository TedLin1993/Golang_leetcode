package leetcode

import "fmt"

func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 1 {
		return heights[0]
	}

	leftIndexes := make([]int, length)
	rightIndexes := make([]int, length)

	leftIndexes[0] = 0
	for i := 1; i < length; i++ {
		leftIdx := i
		for leftIdx-1 >= 0 && heights[leftIdx-1] >= heights[i] {
			leftIdx = leftIndexes[leftIdx-1]
		}
		leftIndexes[i] = leftIdx
	}

	rightIndexes[length-1] = length - 1
	for i := length - 2; i >= 0; i-- {
		rightIdx := i
		for rightIdx+1 < length && heights[rightIdx+1] >= heights[i] {
			rightIdx = rightIndexes[rightIdx+1]
		}
		rightIndexes[i] = rightIdx
	}

	res := 0
	for i := 0; i < length; i++ {
		temp := heights[i] * (rightIndexes[i] - leftIndexes[i] + 1)
		if res < temp {
			res = temp
		}
	}
	return res
}

func largestRectangleArea__(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	dpLeft := make([]int, l)
	dpRight := make([]int, l)

	for i := 1; i < l; i++ {
		dpLeft[i] = i
		j := i
		for dpLeft[j]-1 >= 0 && heights[dpLeft[j]-1] >= heights[i] {
			j = dpLeft[j] - 1
		}
		dpLeft[i] = dpLeft[j]
	}

	dpRight[l-1] = l - 1
	for i := l - 2; i >= 0; i-- {
		dpRight[i] = i
		j := i
		for dpRight[j]+1 <= l-1 && heights[dpRight[j]+1] >= heights[i] {
			j = dpRight[j] + 1
		}
		dpRight[i] = dpRight[j]
	}

	result := 0

	for i := 0; i < l; i++ {
		if heights[i]*(dpRight[i]-dpLeft[i]+1) > result {
			result = heights[i] * (dpRight[i] - dpLeft[i] + 1)
		}
	}
	return result

}

func TestLargestRectangleArea() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}
