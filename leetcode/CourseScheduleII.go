package leetcode

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 1 {
		return []int{0}
	}

	inDegree := make([]int, numCourses)
	edge := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
	}

	result := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			result = append(result, i)
		}
	}

	for i := 0; i < len(result); i++ {
		course := result[i]
		numCourses--
		for j := 0; j < len(edge[course]); j++ {
			inDegree[edge[course][j]]--
			if inDegree[edge[course][j]] == 0 {
				result = append(result, edge[course][j])
			}
		}
	}

	if numCourses != 0 {
		return nil
	}
	return result
}

func TestCourseScheduleII() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{{}}))
	fmt.Println(findOrder(3, [][]int{{1, 0}, {2, 1}, {0, 2}}))
}
