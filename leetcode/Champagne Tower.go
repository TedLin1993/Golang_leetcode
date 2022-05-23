package leetcode

import (
	"fmt"
	"math"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	pouredArray := []float64{float64(poured)}
	for row := 0; row < query_row; row++ {
		columnCount := row + 1
		nextPouredArray := make([]float64, columnCount+1)
		for col := 0; col < columnCount; col++ {
			if pouredArray[col] > 1 {
				overPoured := pouredArray[col] - 1
				nextPouredArray[col] += overPoured / 2
				nextPouredArray[col+1] += overPoured / 2
			}
		}
		pouredArray = nextPouredArray
	}
	return math.Min(1, pouredArray[query_glass])
}

func TestchampagneTower() {
	fmt.Println(champagneTower(1, 1, 1))           //0
	fmt.Println(champagneTower(2, 1, 1))           //0.5
	fmt.Println(champagneTower(100000009, 33, 17)) //1
	fmt.Println(champagneTower(25, 6, 1))          //0.1875
}
