func max(a, b int) int {

	if a > b {
		return a
	} 

	return b
}


func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]

	for i := 1; i < len(nums); i++ {

		dp[i] = max(dp[i-1] + nums[i], nums[i])

		ans = max(ans, dp[i])
	}

	return ans

}