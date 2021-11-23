/*
 * @lc app=leetcode id=153 lang=golang
 */
// @lc code=start
func findMin(nums []int) int {

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return nums[lo]
}
// @lc code=end

