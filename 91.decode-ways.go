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

	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {

		temp := cur
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			cur = pre
		} else if s[i-1] == '1' || s[i-1] == '2' && (s[i] >= '1' && s[i] <= '6') {
			cur = cur + pre
		} else {
			cur = cur
		}
		pre = temp
	}
	return cur
}
// @lc code=end

