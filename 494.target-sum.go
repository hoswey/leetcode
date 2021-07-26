/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {

	//dp[i][j] =  0 if i = 0
	//		   = 1 [+- nums[0]]
	//			  dp[i-1][S+nums[i]] + dp[i-1][S-nums[i]]
	//
	var sum int
	for _, num := range nums {
		sum += num
	}

	if S > sum || S < -sum {
		return 0
	}

	total := sum * 2 + 1
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total)
	}

	dp[0][sum + nums[0]] += 1
	dp[0][sum - nums[0]] += 1

	for i := 1; i < len(nums); i++ {
		for j := -sum; j <= sum; j++ {
			if j + nums[i] > sum {
				dp[i][j+sum] = dp[i-1][j+sum-nums[i]]
			} else if j - nums[i] < -sum {
				dp[i][j+sum] = dp[i-1][j+sum+nums[i]] 
			} else {	
				dp[i][j+sum] = dp[i-1][j+sum+nums[i]] + dp[i-1][j+sum-nums[i]]
			}
		}
	}
    return dp[len(nums)-1][S+sum]
}
// @lc code=end

