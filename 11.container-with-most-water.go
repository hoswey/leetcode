/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {

	var ans int
	left := 0
	right := len(height) - 1

	for left < right {

		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end
