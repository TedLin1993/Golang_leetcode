package leetcode

import (
	"fmt"
)

func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	visit := make([]bool, len(arr))
	valueIndexMap := make(map[int][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		valueIndexMap[arr[i]] = append(valueIndexMap[arr[i]], i)
	}
	numVisit := map[int]bool{}

	visit[0] = true
	layer := 0
	queue := []int{0}
	for len(queue) != 0 {
		layerQueueNum := len(queue)
		for i := 0; i < layerQueueNum; i++ {
			curIndex := queue[i]

			if curIndex == len(arr)-1 {
				return layer
			}

			if curIndex-1 >= 0 && !visit[curIndex-1] {
				visit[curIndex-1] = true
				queue = append(queue, curIndex-1)
			}

			if curIndex+1 < len(arr) && !visit[curIndex+1] {
				visit[curIndex+1] = true
				queue = append(queue, curIndex+1)
			}

			//same number will not for loop twice
			if !numVisit[arr[curIndex]] {
				numVisit[arr[curIndex]] = true
				for _, idx := range valueIndexMap[arr[curIndex]] {
					if !visit[idx] {
						visit[idx] = true
						queue = append(queue, idx)
					}
				}
			}
		}
		queue = queue[layerQueueNum:]
		layer++
	}
	return layer
}

func TestJump_Game_IV() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))                                                                                                                            //3
	fmt.Println(minJumps([]int{7}))                                                                                                                                                                      //0
	fmt.Println(minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))                                                                                                                                                 //1
	fmt.Println(minJumps([]int{6, 1, 9}))                                                                                                                                                                //2
	fmt.Println(minJumps([]int{68, -94, -44, -18, -1, 18, -87, 29, -6, -87, -27, 37, -57, 7, 18, 68, -59, 29, 7, 53, -27, -59, 18, -1, 18, -18, -59, -1, -18, -84, -20, 7, 7, -87, -18, -84, -20, -27})) //5

}
