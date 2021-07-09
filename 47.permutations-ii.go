/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
 
	var ans [][]int
	backtracking(nums, 0, &ans)
	return ans
}

func backtracking(nums []int, start int, ans *[][]int) {

	if start == len(nums) {
		*ans = append(*ans, append([]int(nil), nums...))
		return
	}

	for i := start; i < len(nums); i++ {

		if start != i && nums[start] == nums[i] {
			continue
		}
		
		nums[start], nums[i] = nums[i], nums[start]
		backtracking(nums,start + 1, ans)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
// @lc code=end

