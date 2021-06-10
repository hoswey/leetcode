/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {

	ori := [256]int{}	
	for _, c := range t {
		ori[c]++
	}


	cur := [256]int{}	
	check := func() bool {
		for i, _ := range ori {
			if ori[i] > cur[i] {
				return false
			}
		}
		return true
	}

	ans := s
	var found bool
	var start, end int
	for end < len(s)  {

		cur[s[end]]++

		// fmt.Printf("0. start: %d, end: %d\n", start, end)
		//如何已经包含，则尝试进行收缩start
		if check() {

			found = true
			if (end - start + 1) < len(ans) {
				ans = s[start: end + 1]
			}
			// fmt.Printf("1. start: %d, end: %d, ans:%v\n", start, end, ans)
			for start <= end {
				cur[s[start]]--
				start++					
				//收缩极限
				chk := check()
				if !chk {
					if (end - (start - 1) + 1) < len(ans) {
						ans = s[start - 1: end + 1]
					}
					// fmt.Printf("2. start: %d, end: %d, ans: %v\n", start - 1, end, ans)
					break
				}
			}
		}
		end++
	}

	if !found {
		return ""
	}
	return ans
}
// @lc code=end

