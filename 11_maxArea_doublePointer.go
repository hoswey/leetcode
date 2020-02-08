func maxArea(height []int) int {
    
	var ans int

	i := 0
	j := len(height) - 1

	for i < j {
		ans = max(ans, (j - i) * min(height[i], height[j]))

		if height[i] == height[j] {
			i++
			j--
		} else if height[i] > height[j] {
			j--
		} else {
			i++
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