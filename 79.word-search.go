/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	M := len(board)
	if M == 0 {
		return false
	}
	N := len(board[0])
	visit := make([][]bool, M)
	for i := 0; i < M; i++ {
		visit[i] = make([]bool, N)
	}

	var ans bool
    
	var backtracking func([][]byte, int, int, [][]bool, string, int)
	backtracking = func(board [][]byte, i int, j int, visit [][]bool, word string, pos int) {
		if ans {
			return
		}
		
		if pos == len(word) {
			ans = true
			return
		}

		if i < 0 || i >= M {
			return
		}

		if j < 0 || j >= N {
			return
		}

		if visit[i][j] {
			return
		}
		
		if board[i][j] != word[pos] {
			return
		}
		visit[i][j] = true
		backtracking(board, i - 1, j, visit, word, pos + 1)
		backtracking(board, i + 1, j, visit, word, pos + 1)
		backtracking(board, i, j - 1, visit, word, pos + 1)
		backtracking(board, i, j + 1, visit, word, pos + 1)
		visit[i][j] = false
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			backtracking(board, i, j, visit, word, 0)
		}
	}

	return ans
}
// @lc code=end

