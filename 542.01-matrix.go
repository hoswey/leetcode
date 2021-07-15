/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {

	if len(mat) == 0 {
		return nil
	}

	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				mat[i][j] = (1 << 32) - 1
				if i > 0 {
					mat[i][j] = min(mat[i][j], mat[i-1][j] + 1)
				}
				if j > 0 {
					mat[i][j] = min(mat[i][j], mat[i][j-1] + 1)
				}
			}
		}
	}
    
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if i != m - 1 {
					mat[i][j] = min(mat[i][j], mat[i+1][j] + 1)
				}
				if j != n - 1 {
					mat[i][j] = min(mat[i][j], mat[i][j+1] + 1)
				}
			}
		}
	}
	return mat
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

