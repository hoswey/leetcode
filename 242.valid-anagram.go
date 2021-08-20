/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {

	cnt1 := countChar(s)
	cnt2 := countChar(t)

	for i := 0; i < len(cnt1); i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}

func countChar(s string) []int {

	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	return cnt
}

// @lc code=end

