/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 */

// @lc code=start
func maxCoins(nums []int) int {

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	for l := 3; l <= len(nums); l++ {
		for i := 0; i <= len(nums) - l; i++ {
			j := i + l - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], nums[i] * nums[j] * nums[k] + dp[i][k] + dp[k][j])
			} 
		}
	}
	return dp[0][len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

