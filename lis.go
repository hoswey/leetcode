//[10,9,2,5,3,7,101,18]
//output 4

func lengthOfLIS(nums []int) int {
    
    if len(nums) == 0 {
        return 0
    }

	dp := make([]int, len(nums))
	dp[0] = 1

	for i, _ := range dp {
		dp[i] = 1
	}

    lis := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
                lis = max(dp[i], lis)
			}
		}
	}

	return lis
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}