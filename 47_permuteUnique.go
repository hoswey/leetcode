/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
import "sort"
func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)

	visits := make([]bool, len(nums))
	var ans [][]int
	backtracking(nums, 0, nil, &ans, visits)
	return ans
}

func backtracking(nums []int, idx int, cur []int, ans *[][]int, visits []bool) {

	if idx == len(nums) {
		*ans = append(*ans, cur)
		return
	}

	for i := 0; i < len(nums); i++ {

		if visits[i] || i > 0 && nums[i] == nums[i-1] && !visits[i-1] {
			continue
		}

		visits[i] = true
		arr := make([]int, len(cur))
		copy(arr, cur)
		arr = append(arr, nums[i])
		backtracking(nums, idx+1, arr, ans, visits)
		visits[i] = false
	}
}
// @lc code=end

