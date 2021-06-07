/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {

	ends := make([]int, 26)
	for i, c := range s {
		ends[c - 'a'] = i
	}

	var ans []int
	var start, end int
	for i, c := range s {
		end = max(ends[c - 'a'], end)
		if i == end {
			ans = append(ans, i - start + 1)
			start = i + 1
			end = 0
		}
	}
	return ans
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
// @lc code=end

