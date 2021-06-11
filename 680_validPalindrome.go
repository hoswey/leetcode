/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {

	return doValidPalindrome(s, true);
}

func doValidPalindrome(s string, canRemove bool) bool {

	lo, hi := 0, len(s) - 1
	for lo < hi {
		
		if s[lo] != s[hi] {
			if !canRemove {
				return false
			}
			return doValidPalindrome(s[lo: hi], false) || doValidPalindrome(s[lo+1: hi+1], false)
		}
		lo++
		hi--
	}
	return true

}

// @lc code=end

