var count int

func findTargetSumWays(nums []int, S int) int {

	if len(nums) == 0 {
		return 0
	}

	findTargetSumWays0(nums,0, 0, S)
    return count
}

func findTargetSumWays0(nums []int, i, sum, S int) {

	if len(nums) == i {
		if sum == S {
			count++
		}
		return
	}

	findTargetSumWays0(nums, i+1, sum + nums[i], S)
	findTargetSumWays0(nums, i+1, sum - nums[i], S)
}