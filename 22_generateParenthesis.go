package main

import "fmt"

// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

// 例如，给出 n = 3，生成结果为：

// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
 

func generateParenthesis(n int) []string {

	if n == 0 {
		return nil
	}

	var ans []string
	backtracking(n, n, "", &ans)
	return ans
}

func backtracking(left, right int, path string, ans *[]string) {

	if right < left {
		return
	}
	
	if left == 0 && right == 0 {
		*ans = append(*ans, path)
		return
	} 

	if left > 0 {
		backtracking(left - 1, right, path + "(", ans)
	}

	if right > 0 {
		backtracking(left, right - 1, path + ")", ans)
	}
}

func main() {

	fmt.Printf("%v\n", generateParenthesis(1))
	fmt.Printf("%v\n", generateParenthesis(2))
	fmt.Printf("%v\n", generateParenthesis(3))
}