/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte)  {

	if len(board) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])
	for j:= 0; j < n; j++ {
		backtracking(board, 0, j, m, n)
		backtracking(board, m - 1, j, m, n)
	}

	for i := 0; i < m; i++ {
		backtracking(board, i, 0, m, n)
		backtracking(board, i, n - 1, m, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}
}

func backtracking(board [][]byte, i, j, m, n int) {

	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 'X' || board[i][j] == '-' {
		return
	}

	board[i][j] = '-'

	backtracking(board, i - 1, j, m, n)
	backtracking(board, i, j - 1, m, n)
	backtracking(board, i, j + 1, m, n)
	backtracking(board, i + 1, j, m, n)
}
// @lc code=end

