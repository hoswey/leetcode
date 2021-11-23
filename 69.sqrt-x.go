/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {

	if x <= 1 {
		return x
	}

	low := 1
	high := x / 2

	var ans int
	for low <= high {

		mid := low + (high-low)/2
		p := mid * mid

		if p == x {
			return mid
		} else if p < x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return ans
}
// @lc code=end