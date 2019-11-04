func maxProfit(prices []int) int {

	profits := make([]int, len(prices) + 1)

	for i := 2; i <= len(profits); i++{
		profits[i] = profits[i] - profits[i-1]
	}

	var max int
	for _, profit := range profits {
		if profit > 0 {
			max += profit
		}
	}

	return max
 }