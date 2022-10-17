package leetcode

import "fmt"

var equRoot []int

func equationsPossible(equations []string) bool {
	alphabetCount := 26
	equRoot = make([]int, alphabetCount)
	for i := 0; i < alphabetCount; i++ {
		equRoot[i] = i
	}

	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '=' {
			leftIdx := int(equations[i][0] - 'a')
			rightIdx := int(equations[i][3] - 'a')
			equRoot[findEquations(leftIdx)] = findEquations(rightIdx)
		}
	}

	for i := 0; i < len(equations); i++ {
		leftIdx := int(equations[i][0] - 'a')
		rightIdx := int(equations[i][3] - 'a')
		if equations[i][1] == '!' && equRoot[findEquations(leftIdx)] == findEquations(rightIdx) {
			return false
		}
	}

	return true
}

func findEquations(idx int) int {
	if equRoot[idx] != idx {
		equRoot[idx] = findEquations(equRoot[idx])
	}
	return equRoot[idx]
}

func Test_equationsPossible() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))         //false
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))         //true
	fmt.Println(equationsPossible([]string{"b==a", "c==b", "z==d"})) //true
	fmt.Println(equationsPossible([]string{"a==b", "b!=c", "c==a"})) //false
	fmt.Println(equationsPossible([]string{"a!=a"}))                 //false
}
