/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
import "container/list"

func isValid(s string) bool {

	l := list.New()

	for i := 0; i < len(s); i++ {

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			l.PushFront(s[i])
		} else {
			if l.Len() == 0 {
				return false
			}

			e := l.Front()
			l.Remove(e)
			v := e.Value.(byte)

			if !(s[i] == ')' && v == '(' || s[i] == ']' && v == '[' || s[i] == '}' && v == '{') {
				return false
			}
		}
	}

	return l.Len() == 0
}

// @lc code=end

