/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
    
	var ans [][]int
	comb := make([]int, k)

	var backtracking func(int, int, int)
	backtracking = func(n int, start int, count int) {

		if count == k {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		for i := start; i <= n; i++ {
			comb[count] = i
			count++
			backtracking(n, i + 1, count)			
			count--
		}
	}

	backtracking(n, 1, 0)
	return ans
}
// @lc code=end

