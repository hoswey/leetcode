/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfit(prices []int) int {
    
	// T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	// T[i][k][1] = max(T[i-1][k][1], T[i-1][k-1][0] - prices[i])
	ik0 := make([]int, k + 1)
	ik1 := make([]int, k + 1)
	for k, _ := range ik1 {
		ik1[k] = -(1<<31 - 1)
	}

	for _, price := range prices {
		for j := 1; j <= k; j++ {
			ik0[j] = max(ik0[j], ik1[j] + price)
			ik1[j] = max(ik1[j], ik0[j-2] - price)
		}
	}   
	return ik0[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

