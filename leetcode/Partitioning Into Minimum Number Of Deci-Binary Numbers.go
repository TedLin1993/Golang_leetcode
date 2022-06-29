package leetcode

import "fmt"

func minPartitions(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		res = max(res, int(n[i]-'0'))
	}
	return res
}

func TestminPartitions() {
	fmt.Println(minPartitions("32"))                   //3
	fmt.Println(minPartitions("82734"))                //8
	fmt.Println(minPartitions("27346209830709182346")) //9
}
