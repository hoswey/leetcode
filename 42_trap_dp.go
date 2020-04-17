func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	maxLeft  := make([]int, len(height))
	maxRight := make([]int, len(height))

	for i := 1; i < len(maxLeft); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}

	for i := len(maxRight) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	var ans int
	for i := 1; i < len(height) - 1; i++ {
		m := min(maxLeft[i], maxRight[i])

		if m > height[i] {
			ans += m - height[i]
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}