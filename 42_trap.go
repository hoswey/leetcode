func trap(height []int) int {

	var ans int
	for {

		var set bool
		var li, topHeight int

		for i := 0; i < len(height); i++ {

			if height[i] == 0 {
				continue
			}

			if set {
				ans += (i - li - 1)
				li = i
			} else {
				set = true
				li = i
			}

			height[i] = height[i] - 1
			if height[i] > 0 {
				topHeight++
			}
		}

		if topHeight < 2 {
			break
		}
	}

	return ans
    
}