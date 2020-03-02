func max(a, b int) int {

	if a > b {
		return a
	} 

	return b
}


func maxSubArray(nums []int) int {
    
    if len(nums) == 0 {
    	return 0
    }

    ans := nums[0]
    sum := nums[0]

    for i := 1; i < len(nums); i++ {

    	if sum < 0 {
    		sum = nums[i]
    	} else {
    		sum = sum + nums[i]
    	}

    	ans = max(ans, sum)
    }
}