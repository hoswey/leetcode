//word1 = "horse", word2 = "ros"
//output: 3
// [0 1 2 3]
// [1 1 2 3]
// [2 2 1 2]
// [3 2 2 2]
// [4 3 3 2]
// [5 4 4 3]
func minDistance(word1 string, word2 string) int {

	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1) + 1)

	for i, _ := range dp {
		dp[i] = make([]int, len(word2) + 1)
	}

	dp[0][0] = 0
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for j := 1; j <= len(word2); j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
    
    for i := 1; i <= len(word1); i++ {
    	for j := 1; j <= len(word2); j++ {

    		if word1[i-1] == word2[j-1] {
    			dp[i][j] = dp[i-1][j-1]
    		} else {
    			dp[i][j] = min(min(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j]) + 1
    		}
    	}
    }

    return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a 
	} else {
		return b
	}
}