package main

import "fmt"

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 示例:

// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]

// 给定 word = "ABCCED", 返回 true
// 给定 word = "SEE", 返回 true
// 给定 word = "ABCB", 返回 false
 

func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}


	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtrack(board, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, visited [][]bool, i, j int, word string, k int) bool {

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if visited[i][j] {
		return false
	}

	if word[k] != board[i][j] {
		return false
	}

	if k == len(word) - 1 {
		return true
	}


	visited[i][j] = true

	if backtrack(board, visited, i-1, j, word, k+1) {
		return true
	}

	if backtrack(board, visited, i+1, j, word, k+1) {
		return true
	}

	if backtrack(board, visited, i, j-1, word, k+1) {
		return true
	}

	if backtrack(board, visited, i, j+1, word, k+1) {
		return true
	}		

	visited[i][j] = false

	return false
}

func main() {

	board := [][]byte{
	  {'A','B','C','E'},
	  {'S','F','C','S'},
	  {'A','D','E','E'},
	}
	var word string

	word = "ABCCED"
	fmt.Printf("word:%s, expected:%v, actual:%v\n", word, true, exist(board, word))

	word = "SEE"
	fmt.Printf("word:%s, expected:%v, actual:%v\n", word, true, exist(board, word))

	word = "ABCB"
	fmt.Printf("word:%s, expected:%v, actual:%v\n", word, false, exist(board, word))
}
