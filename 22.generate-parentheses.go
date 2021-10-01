/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {

	var ans []string

	var backtracking func([]byte, int, int)

	backtracking = func(cur []byte, l, r int) {

		if l > n || r > n {
			return
		}

		if l < r {
			return
		}

		if l == n && r == n {
			ans = append(ans, string(cur))
			return
		}

		backtracking(append(cur, '('), l+1, r)
		backtracking(append(cur, ')'), l, r+1)
	}

	backtracking(nil, 0, 0)

	return ans
}

// @lc code=end

