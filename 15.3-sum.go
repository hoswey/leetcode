/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
import "sort"

func threeSum(nums []int) [][]int {
	
	sort.Ints(nums)

	var ans [][]int
	for i := 0; i < len(nums) - 2; i++ {

		j := i + 1
		k := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {

			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return ans
}
// @lc code=end

