/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 */

// @lc code=start
func maxCoins(nums []int) int {
 
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	memo := make([][]int, len(nums))
	for i, _ := range memo {
		memo[i] = make([]int, len(nums))
	}

	var cal func(int, int) int
	cal = func(left, right int) int {

		if right - left < 2 {
			return 0
		}

		if memo[left][right] != 0 {
			return memo[left][right]
		}

		var maxCoins int
		for i := left + 1; i < right; i++ {
			coins := nums[i] * nums[left] * nums[right] + cal(left, i) + cal(i, right)
			maxCoins = max(maxCoins, coins)
		}
		memo[left][right] = maxCoins

		return maxCoins
	}


	return cal(0, len(nums) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

