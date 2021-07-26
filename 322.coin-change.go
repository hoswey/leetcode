/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
    
	dp := make([]int, amount + 1)

	intMax := (1 << 31) - 1
	for i, _ := range dp {
		dp[i] = intMax
	}
	dp[0] = 0
	//dp[i] =  dp[i-nums[k]] + 1  if i >= nums[k]
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i - coin] + 1)
			}
		}
	}
	if dp[amount] == intMax {
		return -1
	}
	return dp[amount]
}

func min(a int, b int)int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

