/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {

	//Input: s = "babad"
	// dp[i][j] =  true   if   i == j
	//          =  true   if  s[i] == s[j] && s[i+1][j-1] = true

	n := len(s)
	if n <= 1 {
		return s
	}

	var ans string
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
		ans = string(s[i])
	}

	for i := 0; i <= n - 2; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans = string(s[i:i+2])
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			j := i + l - 1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				ans = string(s[i:j+1])
			}
		}
	}
	return ans
}
// @lc code=end

