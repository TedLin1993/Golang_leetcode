package leetcode

import "fmt"

type Trie struct {
	Word     string
	Children [26]*Trie
}

var res []string

func findWords(board [][]byte, words []string) []string {
	res = []string{}

	//build trie
	root := &Trie{}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			if node.Children[index] == nil {
				node.Children[index] = &Trie{}
			}
			node = node.Children[index]
		}
		node.Word = word
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			index := board[i][j] - 'a'
			dfsFindWord(board, root.Children[index], i, j)
		}
	}
	return res
}

func dfsFindWord(board [][]byte, node *Trie, row int, column int) {
	if node == nil {
		return
	}
	if node.Word != "" {
		res = append(res, node.Word)
		node.Word = ""
	}

	tempChar := board[row][column]
	board[row][column] = '$'
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for _, dir := range directions {
		nextRow := row + dir[0]
		nextColumn := column + dir[1]
		if nextRow >= 0 && nextRow < len(board) && nextColumn >= 0 && nextColumn < len(board[0]) &&
			board[nextRow][nextColumn] != '$' {
			index := board[nextRow][nextColumn] - 'a'
			dfsFindWord(board, node.Children[index], nextRow, nextColumn)
		}
	}
	board[row][column] = tempChar
}

func TestfindWords() {
	fmt.Println(findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"})) //[eat oath]
	fmt.Println(findWords([][]byte{{'a'}}, []string{"a"}))                                                                                                           //[a]
}
