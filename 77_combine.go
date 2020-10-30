/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {

	visit := make([]bool, n)
	var ans [][]int
	backtracking(0, n,visit, nil, k, &ans)
	return ans
}

func backtracking(from, to int, visit []bool, cur []int, k int, ans *[][]int) {

	if len(cur) == k {
		*ans = append(*ans, cur)
		return
	}

	for i := from; i < to; i++ {

		if visit[i] {
			continue
		}

		newCombine := make([]int, len(cur))
		copy(newCombine, cur)
		visit[i] = true
		newCombine = append(newCombine, i + 1)

		backtracking(i, to, visit, newCombine, k, ans)
		visit[i] = false 
	}
} 
// @lc code=end

