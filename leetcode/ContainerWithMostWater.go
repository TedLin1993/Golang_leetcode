package leetcode

import "fmt"

func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	for left < right {
		area := getAreaAndModifyPosition(height, &left, &right)
		if res < area {
			res = area
		}
	}
	return res
}

func getAreaAndModifyPosition(height []int, left, right *int) int {
	if height[*left] < height[*right] {
		temp := height[*left] * (*right - *left)
		*left++
		return temp
	} else {
		temp := height[*right] * (*right - *left)
		*right--
		return temp
	}
}

func Test_maxArea() {
	var height []int
	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) //49
	height = []int{1, 1}
	fmt.Println(maxArea(height)) //1
	height = []int{4, 3, 2, 1, 4}
	fmt.Println(maxArea(height)) //16
	height = []int{1, 2, 1}
	fmt.Println(maxArea(height)) //2
}
