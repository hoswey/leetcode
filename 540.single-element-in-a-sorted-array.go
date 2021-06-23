/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid % 2 == 0 && nums[mid] == nums[mid + 1] || 
			mid % 2 == 1 && nums[mid] == nums[mid-1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
// @lc code=end