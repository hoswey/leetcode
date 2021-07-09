/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte)  {

	if len(board) == 0 {
		return
	}

	M := len(board)
	N := len(board[0])

	visit := make([][]bool, M)
	for i := 0; i < M; i++ {
		visit[i] = make([]bool, N)
	}

	var backtracking func([][]byte, [][]bool, int, int)
	backtracking = func(board [][]byte, visit [][]bool, i int, j int) {

		if i < 0 || i == M || j < 0 || j == N || visit[i][j] {
			return
		}

		visit[i][j] = true
		if board[i][j] == 'X' {
			return
		}

		board[i][j] = 'Z'

		backtracking(board, visit, i - 1, j)
		backtracking(board, visit, i + 1, j)
		backtracking(board, visit, i, j - 1)
		backtracking(board, visit, i, j + 1)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if i == 0 ||i == M - 1 || j == 0 || j == N - 1 {
				backtracking(board, visit, i, j)
			}
		}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Z' {
				board[i][j] = 'O'
			}

		}
	}
}
// @lc code=end

