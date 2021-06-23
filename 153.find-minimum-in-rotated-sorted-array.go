/*
 * @lc app=leetcode id=153 lang=golang
 */
// @lc code=start
func findMin(nums []int) int {

	lo := 0
	hi := len(nums) - 1

	if nums[lo] < nums[hi] {
		return nums[lo]
	}

	for lo < hi {

		mid := lo + (hi - lo)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[lo] < nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return nums[lo]
}
// @lc code=end

