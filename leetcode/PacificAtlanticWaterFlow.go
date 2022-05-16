package leetcode

func pacificAtlantic(heights [][]int) [][]int {
	length := len(heights)
	width := len(heights[0])

	p := make([][]int, length)
	a := make([][]int, length)
	for i := 0; i < length; i++ {
		p[i] = make([]int, width)
		a[i] = make([]int, width)
	}

	//pacific
	for r := 0; r < length; r++ {
		dfs(r, 0, length, width, &p, heights, 0)
	}
	for c := 0; c < width; c++ {
		dfs(0, c, length, width, &p, heights, 0)
	}

	//atlantic
	for r := 0; r < length; r++ {
		dfs(r, width-1, length, width, &a, heights, 0)
	}
	for c := 0; c < width; c++ {
		dfs(length-1, c, length, width, &a, heights, 0)
	}

	res := [][]int{}
	for r := 0; r < length; r++ {
		for c := 0; c < width; c++ {
			if a[r][c] == 1 && p[r][c] == 1 {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
func dfs(r, c, length, width int, hit *[][]int, heights [][]int, theHeight int) {
	//check valid
	if r < 0 || r >= length || c < 0 || c >= width {
		return
	}
	if (*hit)[r][c] == 1 {
		return
	}

	if heights[r][c] < theHeight {
		return
	}

	(*hit)[r][c] = 1

	dfs(r-1, c, length, width, hit, heights, heights[r][c])
	dfs(r+1, c, length, width, hit, heights, heights[r][c])
	dfs(r, c-1, length, width, hit, heights, heights[r][c])
	dfs(r, c+1, length, width, hit, heights, heights[r][c])
}
