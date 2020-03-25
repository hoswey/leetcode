func change(amount int, coins []int) int {
	return countCoin(coins, len(coins), amount)    
}

func countCoin(coins []int, m int, amount int) int {

	if amount == 0 {
		return 1
	}

	if amount < 0 {
		return 0
	}

	if m <= 0 {
		return 0
	}
	return countCoin(coins, m - 1, amount) + countCoin(coins, m, amount - coins[m-1])
}