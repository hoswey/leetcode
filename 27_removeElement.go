/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := len(nums) - 1

	for i < j {

		for i < j && nums[i] != val {
			i++
		}

		for j > i && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	if nums[i] != val {
		i++
	}
	return i
}
// @lc code=end
