/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 */

// @lc code=start
func minSteps(n int) int {

	dp := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j < i; j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
			}
		}
	}
	return dp[n]
}
// @lc code=end

