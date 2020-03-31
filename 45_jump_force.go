// 给定一个非负整数数组，你最初位于数组的第一个位置。

// 数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。

// 示例:

// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 说明:

// 假设你总是可以到达数组的最后一个位置。
 
func jump(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums))

	return doJump(nums, 0, memo)
}

func doJump(nums []int, position int, memo []int) int {

	if memo[position] != 0 {
		return memo[position]
	}

	if position == len(nums) - 1 {
		return 0
	}

	least := 1 << 31 - 1
	for i := 1; i <= nums[position]; i++ {
		if position + i < len(nums) {
			least = min(least, doJump(nums, position + i, memo) + 1)
			memo[position] = least
		}

	}
	return least
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}