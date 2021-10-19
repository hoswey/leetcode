/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {

	var maxJumps int

	for i := 0; i < len(nums); i++ {

		if i > maxJumps {
			return false
		}

		curJumps := i + nums[i]
		if curJumps > maxJumps {
			maxJumps = curJumps
		}
	}

	return true
}

// @lc code=end

