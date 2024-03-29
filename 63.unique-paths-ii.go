/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}
	
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for j := 1; j < m; j++ {
		for k := 1; k < n; k++ {
			if obstacleGrid[j][k] == 1 {
				dp[j][k] = 0
			} else {
				dp[j][k] = dp[j][k-1] + dp[j-1][k]
			}
		}
	}
	return dp[m-1][n-1]
}
// @lc code=end

