/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
func frequencySort(s string) string {

	m := make(map[byte]int)
	var uniqS string
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = 0
			uniqS += string(s[i])
		}
		m[s[i]]++
	}
	
	str := make([]byte, len(uniqS))
	for i := 0; i < len(uniqS); i++ {
		str[i] = uniqS[i]
	}

	sort.Slice(str, func(i, j int) bool {
		return m[str[i]] > m[str[j]]
	})

	var ans string
	for i := 0; i < len(str); i++ {
		for j := m[str[i]]; j > 0; j-- {
			ans += string(str[i])
		}
	}
	return ans
}
// @lc code=end

