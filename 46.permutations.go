/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {

	var ans [][]int

	var backtracking func([]int, int, int)

	backtracking = func(nums []int, start, end int) {

		if start == end {
			ans = append(ans, append([]int(nil), nums...))
			return
		}

		for i := start; i <= end; i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtracking(nums, start+1, end)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtracking(nums, 0, len(nums) - 1)

	return ans
}

// @lc code=end

