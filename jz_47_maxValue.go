func maxValue(grid [][]int) int {
    
    if len(grid) == 0 {
        return 0
    }
    
    dp := make([][]int, len(grid))
    for i := range dp {
        dp[i] = make([]int, len(grid[0]))
    }
    
    dp[0][0] = grid[0][0]
    
    for col := 1; col < len(grid[0]); col++ {
        dp[0][col] = dp[0][col-1] + grid[0][col] 
    }
    
    for row := 1; row < len(grid); row++ {
        dp[row][0] = dp[row-1][0] + grid[row][0]
    }
    // fmt.Printf("%v\n",dp)
    for row := 1; row < len(grid); row++ {
        for col := 1; col < len(grid[row]); col++ {
            dp[row][col] = max(dp[row-1][col], dp[row][col-1]) + grid[row][col]
        }
    }
    
    // fmt.Printf("%v\n",dp)
    
    return dp[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}