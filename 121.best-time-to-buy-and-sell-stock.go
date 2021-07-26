/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	buy := prices[0]
	var profit int
	for i := 1; i < len(prices); i++ {
		earn := prices[i] - buy
		if earn > profit {
			profit = earn
		} else if prices[i] < buy {
			buy = prices[i]
		}
	}
	return profit
}
// @lc code=end