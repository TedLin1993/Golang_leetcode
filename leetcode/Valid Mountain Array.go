package leetcode

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	//check strict increasing
	if arr[1] <= arr[0] {
		return false
	}
	index := 0
	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if arr[i] < arr[i-1] {
			index = i
			break
		}
	}
	if index == 0 {
		return false
	}

	//check strict decreasing
	for i := index; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}

func TestValidMountainArray() {
	fmt.Println(validMountainArray([]int{2, 1}))       //false
	fmt.Println(validMountainArray([]int{0, 1, 2}))    //false
	fmt.Println(validMountainArray([]int{3, 5, 5}))    //false
	fmt.Println(validMountainArray([]int{0, 3, 2, 1})) //true
}
