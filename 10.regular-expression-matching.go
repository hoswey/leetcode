/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {

	n := len(s)
	m := len(p)

	dp := make([][]bool, n+1)
	for i, _ := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true

	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {

			if p[j-1] != '*' {
				if i != 0 && match(s[i-1], p[j-1]) {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = dp[i][j-2]
				if i != 0 && match(s[i-1], p[j-2]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	return dp[n][m]
}

func match(s, p byte) bool {

	return p == '.' || s == p
}

// @lc code=end
