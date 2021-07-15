package main

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var buy, maxProfit int
	buy = prices[0]

	for i := 1; i < len(prices); i++ {

		if (prices[i] - buy) >= maxProfit {
			maxProfit = prices[i] - buy
		} else if prices[i] < buy {
			buy = prices[i]
		}
	}

	return maxProfit
}
