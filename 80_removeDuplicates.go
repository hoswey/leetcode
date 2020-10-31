/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	cur := nums[0]
	n := 1 
	idx := 1
	for i := 1 ;i < len(nums); i++ {

		if nums[i] != cur {
			nums[idx] = nums[i]
			cur = nums[i]
			idx++
			n = 1
			continue
		}

		if nums[i] == cur && n < 2 {
			n++
			nums[idx] = nums[i]
			idx++

		}
	}
	return idx
}
// @lc code=end

