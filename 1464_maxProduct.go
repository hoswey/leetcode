/*
 * @lc app=leetcode.cn id=1464 lang=golang
 *
 * [1464] 数组中两元素的最大乘积
 */

// @lc code=start
func maxProduct(nums []int) int {

	ans := make([]int, 2)

	for _, num := range nums {
		if num > ans[0] {
			ans[0] = num
		}
		if ans[0] > ans[1] {
			ans[0], ans[1] = ans[1], ans[0]
		}
	}

	return (ans[0] - 1) * (ans[1] - 1)
}
// @lc code=end

