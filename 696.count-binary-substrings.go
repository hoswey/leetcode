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

	var ans int
	for i := 0; i < n-1; i++ {

		if s[i] != s[i+1] {
			ans++
			k := i - 1
			j := i + 2

			for k >= 0 && j < n && s[k] == s[k+1] && s[j] == s[j-1] {
				ans++
				k--
				j++
			}
		}
	}
	return ans
}
// @lc code=end
