/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int)  {

	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	} 
	
	for i := k; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[k], nums[i] = nums[i], nums[k]
			k++	
		}
	}
}
// @lc code=end

