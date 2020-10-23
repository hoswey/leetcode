/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)
	
	for _, str := range strs {

		var count [26]int
		for _, c := range str {
			count[c-'a']++
		}

		var k string
		for _, c := range count {
			k += fmt.Sprintf("%d#",c)
		}

		list, _ := m[k]
		list = append(list, str)
		m[k] = list
	}

	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
// @lc code=end

