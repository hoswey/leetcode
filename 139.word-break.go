/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	//dp[i] = if word in dict then dp[i - len(word)  - i] 
	//		= false
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(wordDict); j++ {

			if dp[i] {
				continue
			}

			word := wordDict[j]
			start := i + 1 - len(word)
			if start >= 0 && string(s[start:start+len(word)]) == word {
				if start == 0 {
					dp[i] = true
				} else {
					dp[i] = dp[start - 1]
				}
			}
		}
	}
	return dp[len(s) - 1]
}
// @lc code=end

