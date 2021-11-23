/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
import "math"

func findPeakElement(nums []int) int {

	lo := 0
	hi := len(nums) - 1

	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt32
		}
		return nums[i]
	}

	for lo < hi {

		mid := lo + (hi-lo)/2

		l := get(mid-1)
		r := get(mid+1)

		c := get(mid)
		if c > l && c > r {
			return mid
		} else if c < r {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// @lc code=end

