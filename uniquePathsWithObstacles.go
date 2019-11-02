func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	if len(obstacleGrid) == 0 {
		return 0
	}

	x := len(obstacleGrid)
	y := len(obstacleGrid[0])

	dp := make([][]int, x)
	for i := range dp {
		dp[i] = make([]int, y)
	}

	for j := y - 1; j >= 0; j-- {

		for i := x - 1; i >= 0; i-- {

			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				//在边缘
				if i == x - 1 && j == y - 1 {
					dp[i][j] = 1 
				} else if i == x - 1 {
					dp[i][j] = dp[i][j+1]
				} else if j == y - 1 {
					dp[i][j] = dp[i+1][j]
				}  else {
					dp[i][j] = dp[i][j+1] + dp[i+1][j]
				}
			}
		}
	}

	return dp[0][0]
}