func longestCommonSubsequence(text1 string, text2 string) int {

	l1 := len(text1)
	l2 := len(text2)

	if l1 == 0 || l2 == 0 {
		return 0
	}  

	dp := make([][]int, l1) 
	for i := range dp {
		dp[i] = make([]int, l2)
	}

	if text1[0] == text2[0] {
		dp[0][0] = 1
	}

	for i := 1; i < l1; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	} 

	for j := 1; j < l2; j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1-1][l2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}