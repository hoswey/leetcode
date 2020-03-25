func change(amount int, coins []int) int {
	
    if amount == 0 {
        return 1
    }
    
	m := len(coins)
	if m == 0 {
		return 0
	}

	dp := make([][]int, amount + 1)

	for i :=  range dp {
		dp[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= amount; i++ {

		for j := 0; j < m; j++ {

			var include, exclude int

			//不含当前金币
			if j > 0 {
				exclude = dp[i][j-1]
			}

			//含当前金币
			if i >= coins[j] {
				include = dp[i-coins[j]][j]
			}

			dp[i][j] = include + exclude
		}
	}

	return dp[amount][m - 1]

}