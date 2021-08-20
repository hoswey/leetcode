/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) int {

	n := len(s)
	if n <= 1 {
		return n
	}

	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}

	var ans int
	for i := 0; i < n; i++ {
		dp[i][i] = true
		ans++
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans++
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			if dp[i+1][i+l-2] && s[i] == s[i+l-1] {
				dp[i][i+l-1] = true
				ans++
			}
		}
	}

	return ans
}

// @lc code=end

