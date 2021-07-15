//152
func maxProduct(nums []int) int {

	//由于存在负数，所以需要维护最小dp, 因为有可能遇到负数相乘直接变到最大。
	//当前最大的遇到负数也会变成最小值，所以遇到负数的时候， min和max需要交换
	if len(nums) == 0 {
		return 0
	}

    ans := nums[0]
    pMax := nums[0]
    pMin := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] < 0 {
			pMax, pMin = pMin, pMax
		}

		pMax = max(nums[i], pMax * nums[i])
		pMin = min(nums[i], pMin * nums[i])

		ans = max(ans, pMax)
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