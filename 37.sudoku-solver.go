/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte)  {

	var row [9][9]bool
	var col [9][9]bool
	var grid [3][3][9]bool

	var spaces [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				idx := int(board[i][j] - '1')
				row[i][idx] = true
				col[j][idx] = true
				grid[i/3][j/3][idx] = true
			}
		}
	}

	var backtracking func(int) bool

	backtracking = func(pos int) bool {

		if pos == len(spaces) {
			return true
		}

		i, j := spaces[pos][0], spaces[pos][1]
		for k := 1; k <= 9; k++ {
			idx := k - 1
			if row[i][idx] || col[j][idx] || grid[i/3][j/3][idx] {
				continue
			}
			row[i][idx] = true
			col[j][idx] = true
			grid[i/3][j/3][idx] = true

			board[i][j] = byte('0' + k)
			if backtracking(pos + 1) {
				return true
			}
			row[i][idx] = false
			col[j][idx] = false
			grid[i/3][j/3][idx] = false
		}
		return false
	}
	backtracking(0)
}
// @lc code=end

