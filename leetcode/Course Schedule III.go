package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })

	courseList := MaxHeap{}
	time := 0
	for _, course := range courses {
		time += course[0]
		heap.Push(&courseList, course[0])
		if time > course[1] {
			time -= heap.Pop(&courseList).(int)
		}
	}
	return courseList.Len()
}

func TestscheduleCourse() {
	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}})) //3
	fmt.Println(scheduleCourse([][]int{{1, 2}}))                                              //1
	fmt.Println(scheduleCourse([][]int{{3, 2}, {4, 3}}))                                      //0
}
