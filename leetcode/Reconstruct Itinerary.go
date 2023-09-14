package leetcode

import "sort"

func findItinerary(tickets [][]string) []string {
	ticketMap := map[string][]string{}
	for _, t := range tickets {
		from, to := t[0], t[1]
		ticketMap[from] = append(ticketMap[from], to)
	}
	for k := range ticketMap {
		sort.Strings(ticketMap[k])
	}

	res := []string{}
	var dfs func(key string)
	dfs = func(key string) {
		if len(ticketMap[key]) == 0 {
			res = append(res, key)
			return
		}
		for len(ticketMap[key]) > 0 {
			next := ticketMap[key][0]
			ticketMap[key] = ticketMap[key][1:]
			dfs(next)
		}
		res = append(res, key)
	}
	startKey := "JFK"
	dfs(startKey)
	return reverseStrings(res)
}

func reverseStrings(strings []string) []string {
	n := len(strings)
	for i := 0; i < len(strings)/2; i++ {
		strings[i], strings[n-1-i] = strings[n-1-i], strings[i]
	}
	return strings
}
