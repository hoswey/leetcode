/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 */

// @lc code=start
func change(amount int, coins []int) int {

	K := len(coins) + 1
	I := amount + 1

	dp := make([][]int, I)
	for i, _ := range dp {
		dp[i] = make([]int, K)
	}

	for i := 0 ; i < K; i++ {
		dp[0][i] = 1
	}

	for k := 1; k < K; k++ {
		for i := 1; i < I; i++ {
			if i >= coins[k-1] {
				dp[i][k] = dp[i][k-1] + dp[i-coins[k-1]][k]
			} else {
				dp[i][k] = dp[i][k-1]
			}
		}
	}
	return dp[I-1][K-1]
}
// @lc code=end