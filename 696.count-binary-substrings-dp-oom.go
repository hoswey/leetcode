/*
 * @lc app=leetcode id=696 lang=golang
 *
 * [696] Count Binary Substrings
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	//00110011
	//dp[i][j] = false  i == j
	//         = ture if dp[i+1][j-1] == true && s[i] == s[i+1] && s[j] == s[j-1]

	n := len(s)
	if len(s) <= 1 {
		return 0
	}

	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}

	var ans int
	for l := 2; l <= n; l += 2 {
		for i := 0; i <= n-l; i++ {
			if l == 2 {
				if s[i] != s[i+1] {
					dp[i][i+1] = true
					ans++
				}
			} else {
				j := i + l - 1
				if s[i] == s[i+1] && s[j] == s[j-1] && dp[i+1][j-1] {
					dp[i][j] = true
					ans++
				}
			}
		}
	}
	return ans
}

// @lc code=end
