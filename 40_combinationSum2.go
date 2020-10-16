/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	//递归，一个数字只有选和不选的可能性
	ans := make(map[string][]int)
	recurse(candidates, 0, target, nil, ans)

	var ret [][]int
	for _, v := range ans {
		ret = append(ret, v)
	}
	return ret
}

func recurse(candidates []int, from int, target int, cur []int, ans map[string][]int) {

	if target == 0 {
		var key string
		for _, v := range cur {
			key = key + "_" + string(v)
		}
		ans[key] = cur
		return
	}

	if from >= len(candidates) {
		return
	}

	if target >= candidates[from] {
		newset := make([]int, len(cur))
		copy(newset, cur)
		newset = append(newset, candidates[from])

		recurse(candidates, from + 1, target - candidates[from], newset, ans)
	}
	recurse(candidates, from + 1, target, cur, ans)
}
// @lc code=end

