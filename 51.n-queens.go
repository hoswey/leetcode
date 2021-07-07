/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {

	if n == 0 {
		return nil
	}

	board := make([][]bool, n)
	for i, _ := range board {
		board[i] = make([]bool, n)
	}

	var ans [][]string
	var backtracking func([][]bool, int) 
	backtracking = func(board [][]bool, i int) {

		if i >= n {
			ans = append(ans, toStringArr(board))
			return
		}

		for j := 0; j < n; j++ {
			if isValid(board, i, j) {
				board[i][j] = true
				backtracking(board, i + 1)
				board[i][j] = false
			}
		}
	}

	backtracking(board, 0)
	return ans
}

func toStringArr(board [][]bool) []string {

	var ret []string
	for i := 0; i < len(board); i++ {
		var row string
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] {
				row += "Q"
			} else {
				row += "."
			}
		}
		ret = append(ret, row)
	}
	return ret
}

func isValid(board [][]bool, i int, j int) bool {

	for  m := 0; m < i; m++ {
		for n := 0; n < len(board[m]); n++ {

			if !board[m][n] {
				continue
			}

			if n == j {
				return false
			}

			if abs(m - i) == abs(n - j) {
				return false
			}
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
// @lc code=end