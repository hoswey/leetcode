/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
	//dp[0] = 0
	//状态转移方程 dp[i] = 1 + min(dp[i - j *j])  j in (1....n)^2 <= i
	dp := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		dp[i] = n
		for j := 1; j * j <= i; j++ {
			dp[i] = min(dp[i], 1 + dp[i - j * j])
		}
	}	
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

