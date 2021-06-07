/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
import "sort"
import "fmt"

func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	var ans [][]int
	var sequence []int
	var dfs func(int, int)

	dfs = func(from int, left int) {

		if left == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}

		if from > len(candidates) {
			return 
		}

		for i := from; i < len(candidates); i++ {

			if left < candidates[i] {
				break
			}

			if i > from && candidates[i] == candidates[i - 1] {
				continue
			}

			sequence = append(sequence, candidates[i])
			dfs(i + 1, left - candidates[i])
			sequence = sequence[:len(sequence) - 1]
		}
	}

	dfs(0, target)
	return ans
}
// @lc code=end

