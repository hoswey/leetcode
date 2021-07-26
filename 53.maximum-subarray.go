/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1] + nums[i])
		ans = max(ans, dp[i])
	}  
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

