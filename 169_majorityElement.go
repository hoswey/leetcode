/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {

	m := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	panic("cannot find majority")
}

// @lc code=end

