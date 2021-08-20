/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
 
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[byte]byte)
	used := make(map[byte]byte)

	for i := 0; i < len(s); i++ {

		v, ok := mapping[s[i]]
		if ok {
			if v != t[i] {
				return false
			}
		} else {
			if _, ok := used[t[i]]; ok {
				return false
			}
			used[t[i]] = t[i]
			mapping[s[i]] = t[i]
		}
	}
	return true
}
// @lc code=end

