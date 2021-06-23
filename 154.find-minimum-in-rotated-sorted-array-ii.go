/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
func findMin(nums []int) int {

	lo := 0
	hi := len(nums) - 1

	if nums[lo] < nums[hi] {
		return nums[lo]
	}

	for lo < hi {
		
		fmt.Printf("lo:%d, hi:%d\n", lo, hi)

		mid := lo + (hi - lo)/2

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[lo] == nums[mid] && nums[mid] == nums[hi]{
			lo++
			hi--
			continue
		}

		if nums[lo] <= nums[mid] && nums[mid] <= nums[hi] {
			return nums[lo]
		}

		if nums[lo] == nums[mid] {
			lo++
			continue
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