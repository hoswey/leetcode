func minPathSum(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	x := len(grid)
	y := len(grid[0])

	dp := make([][]int, x) 

	for i := range dp {
		dp[i] = make([]int, y)
	}


	for j := y - 1; j >= 0; j-- {
		
        for i := x - 1; i >= 0; i-- {           
			if i == x - 1 && j == y - 1 {
				dp[i][j] = grid[i][j]
			} else if j == y - 1 {
				dp[i][j] = grid[i][j] + dp[i + 1][j]
			} else if i == x - 1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[0][0]
}

func min(a, b int) int{

	if a < b {
		return a 
	} else {
		return b
	}
}