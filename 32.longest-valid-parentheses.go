/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
import "container/list"

func longestValidParentheses(s string) int {

	var ans int
	l := list.New()
	l.PushFront(-1)

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			l.PushFront(i)
		} else {
			e := l.Front()
			l.Remove(e)

			e = l.Front()
			if e == nil {
				l.PushFront(i)
			} else {
				size := i - e.Value.(int)
				if size > ans {
					ans = size
				}
			}
		}
	}

	return ans
}

// @lc code=end

