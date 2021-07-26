/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {

	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j * (i -j), j * dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

