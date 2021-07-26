/*
 * @lc app=leetcode id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	
	M := len(word1)
	N := len(word2)

	dp := make([][]int, M + 1)
	for i, _ := range dp {
		dp[i] = make([]int, N+1)
	}

	for m := 0; m <= M; m++ {
		dp[m][0] = m
	}

	for n := 0; n <= N; n++ {
		dp[0][n] = n
	}

	for m := 1; m <= M; m++ {
		for n := 1; n <= N; n++ {
			if word1[m-1] == word2[n-1] {
				dp[m][n] = dp[m-1][n-1]
			} else {
				dp[m][n] = min(dp[m-1][n-1] + 2, min(dp[m-1][n] + 1, dp[m][n-1] + 1))
			}
		}
	}
	return dp[M][N]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

