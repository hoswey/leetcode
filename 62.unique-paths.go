/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {

	if m <= 1 || n <= 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for j := 1; j < m; j++ {
		for k := 1; k < n; k++ {
			dp[j][k] = dp[j][k-1] + dp[j-1][k]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

