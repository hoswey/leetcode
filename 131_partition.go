/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {

	var ans [][]string
	backtracking(s, nil, &ans)
	return ans
}

func backtracking(s string, path []string, ans *[][]string) {

	if len(s) == 0 {
		*ans = append(*ans, path)
		return
	}
	for i := 1; i <= len(s); i++ {
		if isParlindrome(s[0:i]) {
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, s[0:i])
			backtracking(s[i:], newPath, ans)
		}
	}
}

func isParlindrome(s string) bool {

	if len(s) <= 1 {
		return true
	}
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
// @lc code=end

