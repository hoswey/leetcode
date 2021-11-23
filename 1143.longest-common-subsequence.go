/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
 
	//dp[i][j] = dp[i-1][j-1] + 1 (假如text1[i-1] =  text2[j-1])
	//		   = max(dp[i-1][j], dp[i][j-1])
	M := len(text1)
	N := len(text2)

	dp := make([][]int, M + 1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, N + 1)
	}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[M][N]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

