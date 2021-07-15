/*
 * @lc app=leetcode id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {

	if len(nums) <= 2 {
		return 0
	}

	var ans int
	var dp0, dpi int
	for i := 2; i < len(nums); i++ {
		if (nums[i] - nums[i-1]) == (nums[i-1] - nums[i-2]) {
			dpi = dp0 + 1
			ans += dpi
			dp0 = dpi
		} else {
			dp0 = 0
		}
	}
	return ans    
}
// @lc code=end

