package leetcode

import (
	"fmt"
	"sort"
)

type entry struct {
	index    int
	soldiers int
}

func kWeakestRows(mat [][]int, k int) []int {
	count := make([]entry, len(mat), len(mat))
	for r, row := range mat {
		soldiers := 0
		for _, val := range row {
			soldiers += val
		}
		count[r] = entry{r, soldiers}
	}

	sort.Slice(count, func(i, j int) bool {
		si, sj := count[i].soldiers, count[j].soldiers
		return (si == sj && count[i].index < count[j].index) || si < sj
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, count[i].index)
	}

	return res
}

func TestkWeakestRows() {
	// fmt.Println(kWeakestRows([][]int{
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 0},
	// 	{1, 0, 0, 0, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 1}}, 3)) //2,0,3
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}}, 6)) //[4,6,0,1,2,3]
}
