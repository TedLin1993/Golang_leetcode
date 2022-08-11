package leetcode

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func countVowelPermutation(n int) int {
	v := []float64{1, 1, 1, 1, 1}
	vowels := mat.NewDense(5, 1, v)

	data := []float64{}
	data = append(data, []float64{0, 1, 1, 0, 1}...)
	data = append(data, []float64{1, 0, 1, 0, 0}...)
	data = append(data, []float64{0, 1, 0, 1, 0}...)
	data = append(data, []float64{0, 0, 1, 0, 0}...)
	data = append(data, []float64{0, 0, 1, 1, 0}...)

	a := mat.NewDense(5, 5, data)
	a.Pow(a, n-1)

	var m mat.Dense
	m.Mul(a, vowels)
	// fmt.Println(mat.Formatted(&m))
	return int(mat.Sum(m.Grow(0, 5)))
}

func Test_countVowelPermutation() {
	fmt.Println(countVowelPermutation(1)) //5
	fmt.Println(countVowelPermutation(2)) //10
	fmt.Println(countVowelPermutation(5)) //68
}
