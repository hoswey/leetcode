func climbStairs(n int) int {
    
    if n <= 2 {
        return n
    }

	//状态转移公式  dp[i] = dp[i-2] + dp[i-1]
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}