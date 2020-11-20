/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {

	if len(s) == 0 {
		return nil
	}

	dp := calParlindrome(s)
	var ans [][]string
	backtracking(s, nil, &ans, dp, 0)
	return ans
}

func backtracking(s string, path []string, ans *[][]string, dp [][]bool, from int) {

	if len(s) == 0 {
		*ans = append(*ans, path)
		return
	}
	for i := 1; i <= len(s); i++ {
		if dp[from][from+i-1] {
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, s[0:i])
			backtracking(s[i:], newPath, ans, dp,from+i)
		}
	}
}

func calParlindrome(s string) [][]bool {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++{
		dp[i][i] = true
	}

	for i := 0; i < len(s) - 1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for l := 3; l <= len(s); l++ {
		for i := 0; i <= len(s) - l; i++ {
			if s[i] == s[i+l-1] && dp[i+1][i+l-2] {
				dp[i][i+l-1] = true
			}
		}
	}
	return dp
}
// @lc code=end

