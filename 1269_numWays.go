/*
 * @lc app=leetcode.cn id=1269 lang=golang
 *
 * [1269] 停在原地的方案数
 */

// @lc code=start
func numWays(steps int, arrLen int) int {

	dp := make([][]int, steps + 1)

	mod := 1000000007

	arrLen = min(steps + 1, arrLen)
	for i := range dp {
		dp[i] = make([]int, arrLen)
	}

	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j < arrLen; j++ {
			for k := -1; k <= 1; k++ {
				if k + j >= 0 && k + j < arrLen {
					dp[i][j] = (dp[i][j] + dp[i-1][k+j]) % mod
				}
			}
		}
	}
	return dp[steps][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

