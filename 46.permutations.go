/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {

	var ans [][]int

	var backtracking func([]int, int)
	backtracking = func(nums []int, start int) {
		if start == len(nums) - 1 {
			ans = append(ans, append([]int(nil), nums...))
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtracking(nums, start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtracking(nums, 0)
	return ans
}
// @lc code=end

