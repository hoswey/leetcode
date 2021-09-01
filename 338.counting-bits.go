/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}
// @lc code=end

