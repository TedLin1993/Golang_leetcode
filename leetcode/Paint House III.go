package leetcode

// func minCost(houses []int, cost [][]int, m int, n int, target int) int {
// 	dp := make([][][]int, m) //[house][color][neighbor count]
// 	for i := 0; i < m; i++ {
// 		dp[i] = make([][]int, n+1)
// 		for j := 0; j < n+1; j++ {
// 			dp[i][j] = make([]int, target)
// 		}
// 	}

// 	for i := 0; i < m; i++ {
// 		for j := 1; j < n+1; j++ {
// 			for k := 0; k < target; k++ {
// 				if houses[i] != 0 && houses[i] != j {
// 					dp[i][j][k] = math.MaxInt32
// 					continue
// 				}
// 				dp[i][j][k] = min(dp[i-1][j][k-1]+cost[i][j-1], dp[i-1][j][k]+cost[i][j-1])
// 			}
// 		}
// 	}
// }

// func TestminCost() {
// 	fmt.Println(minCost([]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3)) //9
// }
