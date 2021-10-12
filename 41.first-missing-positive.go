/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {

	nums = append([]int{-1}, nums...)
	for i := 1; i < len(nums); i++ {
		for nums[i] != i && nums[i] != -1 {
			if nums[i] >= len(nums) || nums[i] <= 0 {
				nums[i] = -1
			} else {
				if nums[i] == nums[nums[i]] {
					nums[i] = -1
				} else {
					nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
				}
			}
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == -1 {
			return i
		}
	}
	return len(nums)
}
// @lc code=end

