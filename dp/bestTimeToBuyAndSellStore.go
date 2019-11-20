func max(a, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {

	dp := make([][]int, len(prices))
	dp[1] = make([]int, len(prices))
	


	var m int
	for i := 0; i < len(prices) - 1; i++ {
		profit := prices[i+1] - prices[i]
		dp[1][i+1] = profit
		m = max(m, profit)
	}

	for l := 2; l <= len(prices) - 1; l++ {
		for i := 0; i < len(prices) - l; i++ {
			dp[l%2][i+l] = dp[l%2][i+l-1] + dp[l%2][i+l]
			m = max(m, dp[l%2][i+l])
		}
	}
	return m
 }