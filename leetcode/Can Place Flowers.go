package leetcode

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i-1 >= 0 && flowerbed[i-1] == 1 {
			continue
		}
		if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
			continue
		}
		flowerbed[i] = 1
		n--
		if n == 0 {
			return true
		}
	}
	return false
}

func TestCanPlaceFlowers() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
