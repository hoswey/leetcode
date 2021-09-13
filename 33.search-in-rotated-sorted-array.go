/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	//[4,5,6,7,0,1,2]\n3
	lo := 0
	hi := len(nums) - 1

	for lo < hi {

		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		}

		if nums[lo] > nums[mid] && (target >= nums[lo] || target <= nums[mid]) ||
			nums[lo] < nums[mid] && (target >= nums[lo] && target <= nums[mid]) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if nums[lo] == target {
		return lo
	}
	return -1
}

// @lc code=end

