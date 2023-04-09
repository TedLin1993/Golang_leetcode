package leetcode

import "fmt"

func distance(nums []int) []int64 {
	numMap := make(map[int][]int)
	for i, num := range nums {
		numMap[num] = append(numMap[num], i)
	}

	res := make([]int64, len(nums))
	for i, num := range nums {
		if res[i] != 0 {
			continue
		}
		if len(numMap[num]) == 1 {
			res[i] = 0
		}

		distanceSum := sumOfDistanceDifferences(numMap[num])
		for i := 0; i < len(distanceSum); i++ {
			res[numMap[num][i]] = int64(distanceSum[i])
		}
	}
	return res
}

func sumOfDistanceDifferences(nums []int) []int {
	length := len(nums)
	preSum := make([]int, length)
	preSum[0] = 0
	for i := 1; i < length; i++ {
		preSum[i] = preSum[i-1] + nums[i] - nums[0]
	}
	res := make([]int, length)
	res[0] = preSum[length-1]
	for i := 1; i < len(res)-1; i++ {
		rightSum := preSum[length-1] - preSum[i] - (length-i-1)*(nums[i]-nums[0])
		leftSum := (i+1)*(nums[i]-nums[0]) - preSum[i]
		res[i] = leftSum + rightSum
	}
	res[length-1] = (length)*(nums[length-1]-nums[0]) - preSum[length-1]
	return res
}

func Test_distance() {
	// fmt.Println(distance([]int{1, 3, 1, 1, 2}))       //[5,0,3,4,0]
	// fmt.Println(distance([]int{0, 5, 3}))             //[0 0 0]
	// fmt.Println(distance([]int{1, 1, 1, 1, 1}))       //[10 7 6 7 10]
	fmt.Println(distance([]int{2, 0, 2, 2, 6, 5, 2})) //[11,0,7,7,0,0,13]
}
