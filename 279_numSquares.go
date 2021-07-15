func numSquares(n int) int {

	if n == 0 {
		return 0
	}

	dp := make([]int, n + 1)
	
	for i := range dp {
		dp[i] = i
	}

	for i := 2; i <= n; i ++ {
        for j := 2 ; j <= int(math.Sqrt(float64(i))); j++ {
			dp[i] = min(dp[i], dp[i - j * j] + 1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a 
	}

	return b
}