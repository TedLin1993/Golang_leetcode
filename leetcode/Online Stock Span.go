package leetcode

import "fmt"

type StockSpanner struct {
	stack [][2]int
}

func Constructor_StockSpanner() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	res := 1
	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1][0] {
		res += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, res})
	return res
}

func Test_StockSpanner() {
	stockSpanner := Constructor_StockSpanner()
	fmt.Println(stockSpanner.Next(100)) //1
	fmt.Println(stockSpanner.Next(80))  //1
	fmt.Println(stockSpanner.Next(60))  //1
	fmt.Println(stockSpanner.Next(70))  //2
	fmt.Println(stockSpanner.Next(60))  //1
	fmt.Println(stockSpanner.Next(75))  //4
	fmt.Println(stockSpanner.Next(85))  //6

}
