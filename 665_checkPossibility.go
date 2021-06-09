/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {

	var chg bool
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i + 1] {

			if chg {
				return false
			}
			if i == 0 || i > 0 && nums[i-1] <= nums[i + 1] {
				nums[i] = nums[i + 1]
			} else { 
				nums[i + 1] = nums[i]
			}
			chg = true
		}
	}
	return true
}
// @lc code=end

