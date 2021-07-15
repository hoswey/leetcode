/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	//dp[i][j]表示以matrics[i][j]为右下角的最大正方形
	//状态转移方程 dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	side := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			if dp[i][j] > side {
				side = dp[i][j]
			}

		}
	}
	return side * side
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end