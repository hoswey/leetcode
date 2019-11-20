

func rob(nums []int) int {

	if nums == 0 {
		return 0
	}

	if len(nums) == 1{
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]

	if dp[1] < dp[0] {
		dp[1] = dp[0]
	}

	for i := 2; i < len(nums); i++ {

		dp[i] = dp[i-2] + nums[i]

		if dp[i] < dp[i-1] {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(nums)-1]    
}