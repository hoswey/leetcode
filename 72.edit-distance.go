/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {

	//dp[i][j] =  dp[i-1][j-1]  if word1[i] = word2[j] 
	//         = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
	M := len(word1)
	N := len(word2)
	dp := make([][]int, M + 1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, N + 1)
	}

	for j := 0; j <= N; j++ {
		dp[0][j] = j
	}

	for i := 0; i <= M; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
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

