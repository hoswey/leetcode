/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
import "container/list"
import "strings"

func simplifyPath(path string) string {

	l := list.New()
	tokens := strings.Split(path, "/")

	for _, token := range tokens {
		if token == ".." {
			if l.Len() != 0 {
				e := l.Back()
				l.Remove(e)
			}
		} else if token != "/" && token != "." && token != "" {
			l.PushBack(token)
		}
	}

	if l.Len() == 0 {
		return "/"
	}

	var ans string
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		ans += "/" + e.Value.(string)
	}
	return ans
}
// @lc code=end