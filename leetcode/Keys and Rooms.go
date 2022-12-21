package leetcode

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	queue := []int{0}
	for i := 0; i < len(queue); i++ {
		key := rooms[queue[i]]
		for _, room := range key {
			if !visited[room] {
				queue = append(queue, room)
				visited[room] = true
			}
		}
	}
	return len(queue) == len(rooms)
}

func Test_canVisitAllRooms() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))           //true
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}})) //false
}
