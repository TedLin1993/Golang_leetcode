package leetcode

func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+30)
	dp[lastDay] = min(costs[0], min(costs[1], costs[2]))
	dayIdx := len(days) - 2
	for day := lastDay - 1; day >= days[0]; day-- {
		if day > days[dayIdx] {
			dp[day] = dp[day+1]
			continue
		}
		dp[day] = min(costs[0]+dp[day+1], min(costs[1]+dp[day+7], costs[2]+dp[day+30]))
		dayIdx--
	}
	return dp[days[0]]
}
