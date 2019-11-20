func max(a, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {

	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(prices))
	}


	var m int
	for i := 0; i < len(prices) - 1; i++ {
		profit := prices[i+1] - prices[i]
		dp[i][i+1] = profit
		m = max(m, profit)
	}

	for l := 2; l <= len(prices) - 1; l++ {
		for i := 0; i < len(prices) - l; i++ {
			dp[i][i+l] = dp[i][i+l-1] + dp[i+l-1][i+l]
			m = max(m, dp[i][i+l])
		}
	}
	return m
 }