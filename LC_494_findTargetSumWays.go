func findTargetSumWays(nums []int, S int) int {

	if len(nums) == 0 {
		return 0
	}

    fmt.Printf("%v, S:%d\n", nums, S)
	if len(nums) == 1 && (nums[0] == S || nums[0] == -1) {
		return 1
	}

	var ans int

	ans += findTargetSumWays(nums[1:], S + nums[0])
	ans += findTargetSumWays(nums[1:], S - nums[0])
    
    return ans
}