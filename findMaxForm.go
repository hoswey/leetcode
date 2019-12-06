func findMaxForm(strs []string, m int, n int) int {
    
    dp := make([][]int, m + 1)
    for i, _ := range dp {
    	dp[i] = make([]int, n + 1)
    }

    for _, str := range strs {
    	zeros := strings.Count(str, "0")
    	ones  := strings.Count(str, "1")

    	for i := 1; i <= m; i++ {
    		for j := 1; j <= n; j++ {
    			if i >= zeros && j > ones {
    				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones])
    			}
    		}
    	}
    }

    return dp[m][n]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}