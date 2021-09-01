/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	//Input: s = "abcabcbb"
	chars := make(map[byte]bool)
	var left, right int

	var ans int
	for right < len(s) {

		_, ok := chars[s[right]]; 
		for ok && left < right {
			delete(chars, s[left])
			left++
			_, ok = chars[s[right]]
		}

		chars[s[right]] = true
		right++
		ans = max(ans, right-left)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

