/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {


	//dp[i][j] = dp[i-1][j]
	//		   = dp[i-1][j-nums[i]]  //nums[i] < j
	//		   = true  //nums[i] = j
	if len(nums) == 1 {
		return false
	}

	var sum int
	for _, v := range nums {
		sum += v
	}

	if sum & 1 == 1 {
		return false
	}

	target := sum/2
	dp := make([][]bool, len(nums))
	for i := 0 ; i < len(nums); i++ {
		dp[i] = make([]bool, target + 1)
	}

	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] == target {
				dp[i][j] = true
			}
			if nums[i] < j {
				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
			}
		}
	}
    return dp[len(nums)-1][target]
}
// @lc code=end

