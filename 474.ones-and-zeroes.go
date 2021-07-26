/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n + 1)
	}

	for _, str := range strs {

		one, zero := countOneAndZero(str)
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i - zero][j - one] + 1)
			}
		}

		fmt.Printf("str:%s, dp:%v\n", str, dp)
	} 
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func countOneAndZero(s string) (int, int) {

	var one, zero int
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		} else {
			zero++
		}
	}
	return one, zero
}
// @lc code=end