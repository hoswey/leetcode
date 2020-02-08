func maxArea(height []int) int {
    
    var ans int
	for i := 0; i < len(height) - 1; i++ {
		for j := i + 1; j < len(height); j++ {
            ans = max(ans, (j - i) * min(height[i], height[j]))
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}