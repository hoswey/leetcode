/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {

	maxJumps := nums[0]
	i := 1

	var step int
	for i < len(nums) {
		var nextJumps int
		for i <= maxJumps && i < len(nums) {
			nextJumps = max(nextJumps, i+nums[i])
			i++
		}
		
		step++
		maxJumps = nextJumps
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

