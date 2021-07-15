/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {

	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	//dp[-1], dp[0] = 1
	pre, cur := 1, 1 
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
			cur = pre + cur
		} else {
			cur = cur
		}
		pre = tmp
	}
	return cur
}
// @lc code=end

