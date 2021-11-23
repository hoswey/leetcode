/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	//Input: nums = [5,7,7,8,8,10], target = 8
	//Output: [3,4]

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	left := lo

	if nums[left] != target {
		return []int{-1, -1}
	}

	lo = 0
	hi = len(nums) - 1

	for lo < hi {
		mid := lo + (hi-lo)/2 + 1
		if nums[mid] <= target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	right := lo
	if nums[right] != target {
		return []int{-1, -1}
	}

	return []int{left, right}
}

// @lc code=end

