func maxProfit(prices []int, fee int) int {

	if len(prices <= 1) {
		return 0
	}

	hold := make([]int, len(prices))
	unhold := make([]int, len(prices))

	hold[0] = -prices[0]

	for i := 1; i < len(prices); i++ {

		hold[i] = max(hold[i-1], unhold[i-1] - prices[i])
		unhold[i] = max(unhold[i-1], hold[i-1] + prices[i] - fee)
	}

	return unhold[len(prices) - 1]

}

func max(a, b int) int {

	if a >= b { 
		return a
	} else {
		return b
	}
}