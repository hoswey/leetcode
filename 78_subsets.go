/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {

	var ans [][]int
	backtracking(nums, 0, nil, &ans)
	return ans

}

func backtracking(nums[] int, from int, cur []int, ans *[][]int) {

	if from >= len(nums) {
		*ans = append(*ans, cur)
		return 
	}

	//选
	dst := make([]int, len(cur))
	copy(dst, cur)
	dst = append(dst, nums[from])
	backtracking(nums, from + 1, dst, ans)

	//不选
	backtracking(nums, from + 1, cur, ans)
}

// @lc code=end

