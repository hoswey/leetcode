/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	//T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	//T[i][k][1] = max(T[i-1][k][1], T[i-2][k - 1][0] - prices[i])

	//T[i][0] = max(T[i-1][0], T[i-1][1] + prices[i])
	//T[i][1] = max(T[i-1][1], T[i-2][0] - prices[i])
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		if i >= 2 {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}
	}
	return dp[n-1][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

