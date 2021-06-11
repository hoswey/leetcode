/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {

	start := 0
	end := int(math.Sqrt(float64(c)))

	for start <= end {
		sum := start * start + end * end
		if sum == c {
			return true
		} else if sum < c {
			start++
		} else {
			end--
		}
	}
	return false
}
// @lc code=end

