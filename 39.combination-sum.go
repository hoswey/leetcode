/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
    
	sort.Ints(candidates)
	var ans [][]int
	backtracking(candidates, 0, target, nil,&ans)
	return ans
}

func backtracking(candidates []int, from int, target int, path []int, ans *[][]int) {

	if target == 0 {
		*ans = append(*ans, append([]int(nil), path...))
		return
	}

	if from == len(candidates) {
		return
	}

	if candidates[from] <= target {
		//选
		backtracking(candidates, from, target - candidates[from], append(path, candidates[from]), ans)
		//不选
		backtracking(candidates, from + 1, target, path, ans)
	}
}
// @lc code=end

