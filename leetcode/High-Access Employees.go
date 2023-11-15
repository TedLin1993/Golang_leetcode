package leetcode

import (
	"sort"
	"strconv"
)

func findHighAccessEmployees(access_times [][]string) []string {
	accessMap := make(map[string][]int, 0)
	for _, a := range access_times {
		key := a[0]
		time, _ := strconv.Atoi(a[1])
		accessMap[key] = append(accessMap[key], time)
	}
	res := make([]string, 0, len(accessMap))
	for k, v := range accessMap {
		sort.Ints(v)
		for i := 2; i < len(v); i++ {
			if v[i]-v[i-2] < 100 {
				res = append(res, k)
				break
			}
		}
	}
	return res
}
