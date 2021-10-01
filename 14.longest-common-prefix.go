/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]

	for i := 1; i < len(strs); i++ {
		ans = lcp(ans, strs[i])
	}
	return ans
}

func lcp(str1, str2 string) string {

	var i int
	for ; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			break
		}
	}
	return str1[:i]
}
// @lc code=end

