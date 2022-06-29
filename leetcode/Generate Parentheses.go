package leetcode

import "fmt"

var generateParenthesisResult []string

func generateParenthesis(n int) []string {
	generateParenthesisResult = []string{}
	dfs_generateParenthesis("", 0, 0, n)
	return generateParenthesisResult
}

func dfs_generateParenthesis(str string, open int, close int, n int) {
	if len(str) == n*2 {
		generateParenthesisResult = append(generateParenthesisResult, str)
		return
	}

	if open < n {
		dfs_generateParenthesis(str+"(", open+1, close, n)
	}
	if close < open {
		dfs_generateParenthesis(str+")", open, close+1, n)
	}
}

func TestgenerateParenthesis() {
	fmt.Println(generateParenthesis(1)) //[()]
	fmt.Println(generateParenthesis(2)) //[(()) ()()]
	fmt.Println(generateParenthesis(3)) //[((())) (()()) ()(()) (())() ()()()]
	fmt.Println(generateParenthesis(4)) //["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
}
