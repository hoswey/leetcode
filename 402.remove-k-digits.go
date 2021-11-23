/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
import "container/list"

func removeKdigits(num string, k int) string {

	l := list.New()

	remain := len(num) - k
	for i := 0; i < len(num); i++ {

		for k > 0 && l.Len() > 0 && l.Back().Value.(byte) > num[i] {
			l.Remove(l.Back())
			k--
		}
		l.PushBack(num[i])
	}

	var ans []byte
	leadingZero := true
	for remain > 0 {
		e := l.Front()
		l.Remove(e)

		remain--
		if leadingZero && e.Value.(byte) == '0' {
			continue
		}
		leadingZero = false
		ans = append(ans, e.Value.(byte))
	}

	if ans == nil {
		return "0"
	}
	return string(ans)
}
// @lc code=end

