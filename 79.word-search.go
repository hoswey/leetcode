/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	M := len(board)
	N := len(board[0])

	visits := make([][]bool, M)
	for i, _ := range visits {
		visits[i] = make([]bool, N)
	}

	var backtracking func([][]byte, [][]bool, int, int, int) bool

	backtracking = func(board [][]byte, visits [][]bool, i, j, k int) bool {

		if k == len(word) {
			return true
		}

		if i < 0 || i >= M || j < 0 || j >= N {
			return false
		}

		if visits[i][j] {
			return false
		}

		if board[i][j] != word[k] {
			return false
		}

		visits[i][j] = true

		if backtracking(board, visits, i-1, j, k+1) {
			return true
		}

		if backtracking(board, visits, i+1, j, k+1) {
			return true
		}

		if backtracking(board, visits, i, j-1, k+1) {
			return true
		}

		if backtracking(board, visits, i, j+1, k+1) {
			return true
		}

		visits[i][j] = false
		return false
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if backtracking(board, visits, i, j, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end

