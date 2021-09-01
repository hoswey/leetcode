/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {

	n := len(height)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	var ans int
	for i := 0; i < n; i++ {
		ans += (min(maxLeft[i], maxRight[i]) - height[i])
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

