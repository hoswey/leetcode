/*
 * @lc app=leetcode id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 */

// @lc code=start
import "container/list"

func removeDuplicateLetters(s string) string {

	remain := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := remain[s[i]]; ok {
			remain[s[i]] = remain[s[i]] + 1
		} else {
			remain[s[i]] = 1
		}
	}

	l := list.New()
	exists := make(map[byte]bool)
	for i := 0; i < len(s); i++ {

		if exist, ok := exists[s[i]]; !ok || !exist {
			for l.Len() > 0 && l.Back().Value.(byte) > s[i] && remain[l.Back().Value.(byte)] > 0 {
				e := l.Back()
				l.Remove(e)
				exists[e.Value.(byte)] = false
			}
			exists[s[i]] = true
			l.PushBack(s[i])
		}
		remain[s[i]] = remain[s[i]] - 1
	}

	var ans []byte
	for e := l.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value.(byte))
	}
	return string(ans)
}
// @lc code=end

