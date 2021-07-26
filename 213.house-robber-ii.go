/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {

	//dp[i] = nums[0]			    		  if i == 0
	//		= max(nums[0], nums[1]) 		  if i == 1
    //      = max(dp[i-2] + nums[i], dp[i-1]) if i >= 2
	//		= 偶数 i-2, 奇数 i-1			   if i == len(nums) - 1
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(doRob(nums[0:len(nums)-1]), doRob(nums[1:]))
}

func doRob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[len(nums) - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end