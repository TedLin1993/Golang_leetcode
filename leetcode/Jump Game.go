package leetcode

import "fmt"

func canJump(nums []int) bool {
	vistedLastIndex := 0
	queue := []int{0}
	for len(queue) > 0 {
		index := queue[0]

		//can reach the last index
		if index+nums[index] >= len(nums)-1 {
			return true
		}

		if index+nums[index] > vistedLastIndex {
			for i := vistedLastIndex + 1; i <= index+nums[index]; i++ {
				queue = append(queue, i)
			}
			vistedLastIndex = index + nums[index]
		}
		queue = queue[1:]
	}
	return false
}

func Test_canJump() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) //true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) //false
	fmt.Println(canJump([]int{1, 1, 0, 1}))    //false
}
